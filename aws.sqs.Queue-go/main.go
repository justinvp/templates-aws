package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/sqs"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"deadLetterTargetArn": aws_sqs_queue.Queue_deadletter.Arn,
			"maxReceiveCount":     4,
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err := sqs.NewQueue(ctx, "queue", &sqs.QueueArgs{
			DelaySeconds:            pulumi.Int(90),
			MaxMessageSize:          pulumi.Int(2048),
			MessageRetentionSeconds: pulumi.Int(86400),
			ReceiveWaitTimeSeconds:  pulumi.Int(10),
			RedrivePolicy:           pulumi.String(json0),
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("production"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

