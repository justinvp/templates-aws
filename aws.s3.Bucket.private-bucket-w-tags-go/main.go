package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := s3.NewBucket(ctx, "bucket", &s3.BucketArgs{
			Acl: pulumi.String("private"),
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("Dev"),
				"Name":        pulumi.String("My bucket"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

