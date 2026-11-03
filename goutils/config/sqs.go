package config

import "fmt"

type SqsConfigStaticI interface {
	GetPrefix() string
	GetStackName() string
	ValidateSqsConfig() error
	GetFromQueue() string
	GetToQueue() string
}

type SqsConfigStatic struct {
	StackName string `conf:"env:AWS_RESOURCE_CFN_STACK_NAME"`
	Prefix    string `conf:"env:AWS_RESOURCE_PREFIX"`
	FromQueue string `conf:"env:AWS_SQS_RESULT_CHAN"`
	ToQueue   string `conf:"env:AWS_SQS_FOR_PROCESSING_CHAN"`
}

func (c *SqsConfigStatic) GetPrefix() string {
	return c.Prefix
}

func (c *SqsConfigStatic) GetStackName() string {
	return c.StackName
}

func (c *SqsConfigStatic) ValidateSqsConfig() error {
	if c.StackName == "" {
		return fmt.Errorf("stack name is required")
	}
	if c.Prefix == "" {
		return fmt.Errorf("prefix is required")
	}
	return nil
}

func (c *SqsConfigStatic) GetFromQueue() string {
	return c.FromQueue
}

func (c *SqsConfigStatic) GetToQueue() string {
	return c.ToQueue
}
