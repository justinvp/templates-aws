package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleBucket, err := s3.NewBucket(ctx, "exampleBucket", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewAccessPoint(ctx, "exampleAccessPoint", &s3.AccessPointArgs{
			Bucket: exampleBucket.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

