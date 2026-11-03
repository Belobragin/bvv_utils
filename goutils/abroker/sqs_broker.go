package abroker

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	sqstypes "github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/belobragin/bvv_utils/goutils/awsutil"
	"github.com/belobragin/bvv_utils/goutils/mistake"
	"github.com/belobragin/bvv_utils/goutils/util"
	"go.uber.org/zap"
)

const (
	stackCheckTimeout    = 15 * time.Second
	visibilityTimeoutEnv = "SQS_VISIBILITY_TIMEOUT"
	receiveWaitSeconds   = 20
	receiveMaxMessages   = 10
	receiveRetryDelay    = 5 * time.Second
)

// SqsMsg adapts an sqs message to AllMessageI.
type SqsMsg struct {
	msg   sqstypes.Message
	queue string
}

func (sm *SqsMsg) GetData() []byte {
	return []byte(aws.ToString(sm.msg.Body))
}

// GetKey returns the receipt handle.
func (sm *SqsMsg) GetKey() string {
	return aws.ToString(sm.msg.ReceiptHandle)
}

// GetSubject returns the url of the queue the message came from.
func (sm *SqsMsg) GetSubject() string {
	return sm.queue
}

/*

resource_stack_name = os.environ["AWS_RESOURCE_CFN_WEREWOLF_STACK_NAME"]
        bucket_output_key = f"{os.environ['AWS_RESOURCE_PREFIX']}ProjectBucketName"
        self.image_storage_path = os.environ['IMAGE_CLOUD_STORAGE']
        self.adversarial_storage_path = os.environ['ADVERSARIAL_CLOUD_STORAGE']
        self.filter_storage_path = os.environ['FILTER_CLOUD_STORAGE']
        self.bucket_name = _bucket_name(resource_stack_name, bucket_output_key)
        self.project_bucket = boto3.resource('s3').Bucket(self.bucket_name)

*/

type SqsBroker struct {
	inKey             string
	outKey            string
	visibilityTimeout int
	sqs               *sqs.Client
	queueFrom         string // queue url to recieve messages
	queueTo           string // queue url to send messages
	zapstruct         *zap.Logger
	// InputMsgFoo processes received messages; the message is deleted from the queue only if it returns nil.
	InputMsgFoo func(AllMessageI) error
}

// initialize sqs broker structure
func NewSqsBroker(prefix, stackName string,
	fromQueue, toQueue string,
	foo func(AllMessageI) error,
	w *zap.Logger) (*SqsBroker, error) {
	var (
		sb  = new(SqsBroker)
		err error
	)
	sb.zapstruct = w
	sb.InputMsgFoo = foo

	if len(prefix) == 0 || len(stackName) == 0 {
		return nil, mistake.ErrSubscribeData
	}

	ctx, cancel := context.WithTimeout(context.Background(), stackCheckTimeout)
	defer cancel()

	if err := awsutil.CheckStackExists(ctx, stackName); err != nil {
		return nil, err
	}

	sb.inKey = fmt.Sprintf("%s%s", prefix, fromQueue)
	sb.outKey = fmt.Sprintf("%s%s", prefix, toQueue)
	if sb.visibilityTimeout, err = util.VisibilityTimeoutFromEnv(visibilityTimeoutEnv); err != nil {
		return nil, err
	}
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("load aws config: %w", err)
	}
	sb.sqs = sqs.NewFromConfig(cfg)

	if sb.queueFrom, err = sb.queueURL(ctx, cfg, stackName, sb.inKey); err != nil {
		return nil, err
	}
	if sb.queueTo, err = sb.queueURL(ctx, cfg, stackName, sb.outKey); err != nil {
		return nil, err
	}

	return sb, nil
}

// queueURL resolves the queue name from the stack output and returns its url.
func (sb *SqsBroker) queueURL(ctx context.Context, cfg aws.Config, stackName, key string) (string, error) {
	name, err := awsutil.GetStackOutput(ctx, cfg, stackName, key)
	if err != nil {
		return "", err
	}

	out, err := sb.sqs.GetQueueUrl(ctx, &sqs.GetQueueUrlInput{QueueName: aws.String(name)})
	if err != nil {
		return "", fmt.Errorf("get queue url %s: %w", name, err)
	}

	return aws.ToString(out.QueueUrl), nil
}

func (sb *SqsBroker) GetSqsClient() *sqs.Client {
	return sb.sqs
}

func (sb *SqsBroker) GetQueueFrom() string {
	return sb.queueFrom
}

func (sb *SqsBroker) GetQueueTo() string {
	return sb.queueTo
}

// GetPublishSubject returns the url of the queue events are published to.
func (sb *SqsBroker) GetPublishSubject() string {
	return sb.queueTo
}

func (sb *SqsBroker) PublishEvent(data []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), stackCheckTimeout)
	defer cancel()

	_, err := sb.sqs.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    aws.String(sb.queueTo),
		MessageBody: aws.String(string(data)),
	})
	if err != nil {
		return fmt.Errorf("send message to %s: %w", sb.queueTo, err)
	}
	return nil
}

// ListenEvent long-polls queueFrom until a value arrives on sch; receive errors are logged, sent to errC without blocking, and retried.
func (sb *SqsBroker) ListenEvent(errC chan error, sch chan struct{}) {
	if sb.InputMsgFoo == nil {
		errC <- mistake.ErrSubscribeData
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		select {
		case <-sch:
			cancel()
		case <-ctx.Done():
		}
	}()

	for {
		in := &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(sb.queueFrom),
			MaxNumberOfMessages: receiveMaxMessages,
			WaitTimeSeconds:     receiveWaitSeconds,
		}
		if sb.visibilityTimeout > 0 {
			in.VisibilityTimeout = int32(sb.visibilityTimeout)
		}
		out, err := sb.sqs.ReceiveMessage(ctx, in)
		if err != nil {
			if ctx.Err() != nil {
				return
			}
			err = fmt.Errorf("receive message from %s: %w", sb.queueFrom, err)
			sb.zapstruct.Error("receive sqs message", zap.Error(err))
			select {
			case errC <- err:
			default:
			}
			select {
			case <-time.After(receiveRetryDelay):
				continue
			case <-ctx.Done():
				return
			}
		}

		for _, m := range out.Messages {
			if err := sb.InputMsgFoo(&SqsMsg{msg: m, queue: sb.queueFrom}); err != nil {
				sb.zapstruct.Error("process sqs message", zap.Error(err))
				continue
			}
			_, err := sb.sqs.DeleteMessage(ctx, &sqs.DeleteMessageInput{
				QueueUrl:      aws.String(sb.queueFrom),
				ReceiptHandle: m.ReceiptHandle,
			})
			if err != nil {
				sb.zapstruct.Error("delete sqs message", zap.Error(err))
			}
		}
	}
}

// CloseBroker is a no-op: the sqs client holds no persistent connection.
func (sb *SqsBroker) CloseBroker() error {
	return nil
}

var _ AllBrokerI = (*SqsBroker)(nil)

func SimpleInputMsgFoo(w *zap.Logger) func(AllMessageI) error {
	return func(msg AllMessageI) error {
		w.Info("received message", zap.Any("msg", msg))
		return nil
	}
}
