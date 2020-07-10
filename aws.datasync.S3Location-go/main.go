package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/datasync"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datasync.NewS3Location(ctx, "example", &datasync.S3LocationArgs{
			S3BucketArn: pulumi.Any(aws_s3_bucket.Example.Arn),
			S3Config: &datasync.S3LocationS3ConfigArgs{
				BucketAccessRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
			},
			Subdirectory: pulumi.String("/example/prefix"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

