package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/sqs"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		queue, err := sqs.NewQueue(ctx, "queue", nil)
		if err != nil {
			return err
		}
		_, err = sqs.NewQueuePolicy(ctx, "test", &sqs.QueuePolicyArgs{
			Policy: queue.Arn.ApplyT(func(arn string) (string, error) {
				return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Id\": \"sqspolicy\",\n", "  \"Statement\": [\n", "    {\n", "      \"Sid\": \"First\",\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": \"*\",\n", "      \"Action\": \"sqs:SendMessage\",\n", "      \"Resource\": \"", arn, "\",\n", "      \"Condition\": {\n", "        \"ArnEquals\": {\n", "          \"aws:SourceArn\": \"", aws_sns_topic.Example.Arn, "\"\n", "        }\n", "      }\n", "    }\n", "  ]\n", "}\n", "\n"), nil
			}).(pulumi.StringOutput),
			QueueUrl: queue.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

