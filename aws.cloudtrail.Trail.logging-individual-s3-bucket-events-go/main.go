package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudtrail"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		important_bucket, err := s3.LookupBucket(ctx, &s3.LookupBucketArgs{
			Bucket: "important-bucket",
		}, nil)
		if err != nil {
			return err
		}
		_, err = cloudtrail.NewTrail(ctx, "example", &cloudtrail.TrailArgs{
			EventSelectors: cloudtrail.TrailEventSelectorArray{
				&cloudtrail.TrailEventSelectorArgs{
					DataResource: pulumi.MapArray{
						pulumi.Map{
							"type": pulumi.String("AWS::S3::Object"),
							"values": pulumi.StringArray{
								pulumi.String(fmt.Sprintf("%v%v", important_bucket.Arn, "/")),
							},
						},
					},
					IncludeManagementEvents: pulumi.Bool(true),
					ReadWriteType:           pulumi.String("All"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

