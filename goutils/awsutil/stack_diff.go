package awsutil

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/smithy-go"
	"github.com/belobragin/bvv_utils/goutils/mistake"
)

// checkStackExists returns mistake.ErrAwsStackNotFound if the cfn stack is absent.
func CheckStackExists(ctx context.Context, stackName string) error {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return fmt.Errorf("load aws config: %w", err)
	}

	_, err = cloudformation.NewFromConfig(cfg).DescribeStacks(ctx,
		&cloudformation.DescribeStacksInput{StackName: aws.String(stackName)})
	if err != nil {
		// DescribeStacks reports a missing stack as a ValidationError.
		var apiErr smithy.APIError
		if errors.As(err, &apiErr) && apiErr.ErrorCode() == "ValidationError" &&
			strings.Contains(apiErr.ErrorMessage(), "does not exist") {
			return fmt.Errorf("%w: %s", mistake.ErrAwsStackNotFound, stackName)
		}
		return fmt.Errorf("describe aws stack %s: %w", stackName, err)
	}

	return nil
}

// GetStackOutput returns the value of the cfn stack output with the given key.
func GetStackOutput(ctx context.Context, cfg aws.Config, stackName, outputKey string) (string, error) {
	out, err := cloudformation.NewFromConfig(cfg).DescribeStacks(ctx,
		&cloudformation.DescribeStacksInput{StackName: aws.String(stackName)})
	if err != nil {
		return "", fmt.Errorf("describe aws stack %s: %w", stackName, err)
	}

	for _, st := range out.Stacks {
		for _, o := range st.Outputs {
			if aws.ToString(o.OutputKey) == outputKey {
				return aws.ToString(o.OutputValue), nil
			}
		}
	}

	return "", fmt.Errorf("%w: %s in stack %s", mistake.ErrAwsStackOutputNotFound, outputKey, stackName)
}
