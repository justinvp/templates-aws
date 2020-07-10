package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elastictranscoder"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elastictranscoder.NewPipeline(ctx, "bar", &elastictranscoder.PipelineArgs{
			ContentConfig: &elastictranscoder.PipelineContentConfigArgs{
				Bucket:       pulumi.Any(aws_s3_bucket.Content_bucket.Bucket),
				StorageClass: pulumi.String("Standard"),
			},
			InputBucket: pulumi.Any(aws_s3_bucket.Input_bucket.Bucket),
			Role:        pulumi.Any(aws_iam_role.Test_role.Arn),
			ThumbnailConfig: &elastictranscoder.PipelineThumbnailConfigArgs{
				Bucket:       pulumi.Any(aws_s3_bucket.Thumb_bucket.Bucket),
				StorageClass: pulumi.String("Standard"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

