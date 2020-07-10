package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		bucket, err := s3.NewBucket(ctx, "bucket", nil)
		if err != nil {
			return err
		}
		topic, err := sns.NewTopic(ctx, "topic", &sns.TopicArgs{
			Policy: bucket.Arn.ApplyT(func(arn string) (string, error) {
				return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "    \"Version\":\"2012-10-17\",\n", "    \"Statement\":[{\n", "        \"Effect\": \"Allow\",\n", "        \"Principal\": {\"AWS\":\"*\"},\n", "        \"Action\": \"SNS:Publish\",\n", "        \"Resource\": \"arn:aws:sns:*:*:s3-event-notification-topic\",\n", "        \"Condition\":{\n", "            \"ArnLike\":{\"aws:SourceArn\":\"", arn, "\"}\n", "        }\n", "    }]\n", "}\n", "\n"), nil
			}).(pulumi.StringOutput),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketNotification(ctx, "bucketNotification", &s3.BucketNotificationArgs{
			Bucket: bucket.ID(),
			Topics: s3.BucketNotificationTopicArray{
				&s3.BucketNotificationTopicArgs{
					Events: pulumi.StringArray{
						pulumi.String("s3:ObjectCreated:*"),
					},
					FilterSuffix: pulumi.String(".log"),
					TopicArn:     topic.Arn,
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

