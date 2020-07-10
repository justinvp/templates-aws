package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/sqs"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		bucket, err := s3.NewBucket(ctx, "bucket", nil)
		if err != nil {
			return err
		}
		queue, err := sqs.NewQueue(ctx, "queue", &sqs.QueueArgs{
			Policy: bucket.Arn.ApplyT(func(arn string) (string, error) {
				return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": \"*\",\n", "      \"Action\": \"sqs:SendMessage\",\n", "	  \"Resource\": \"arn:aws:sqs:*:*:s3-event-notification-queue\",\n", "      \"Condition\": {\n", "        \"ArnEquals\": { \"aws:SourceArn\": \"", arn, "\" }\n", "      }\n", "    }\n", "  ]\n", "}\n", "\n"), nil
			}).(pulumi.StringOutput),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketNotification(ctx, "bucketNotification", &s3.BucketNotificationArgs{
			Bucket: bucket.ID(),
			Queues: s3.BucketNotificationQueueArray{
				&s3.BucketNotificationQueueArgs{
					Events: pulumi.StringArray{
						pulumi.String("s3:ObjectCreated:*"),
					},
					FilterPrefix: pulumi.String("images/"),
					Id:           pulumi.String("image-upload-event"),
					QueueArn:     queue.Arn,
				},
				&s3.BucketNotificationQueueArgs{
					Events: pulumi.StringArray{
						pulumi.String("s3:ObjectCreated:*"),
					},
					FilterPrefix: pulumi.String("videos/"),
					Id:           pulumi.String("video-upload-event"),
					QueueArn:     queue.Arn,
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

