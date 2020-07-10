package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		examplebucket, err := s3.NewBucket(ctx, "examplebucket", &s3.BucketArgs{
			Acl: pulumi.String("private"),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketObject(ctx, "examplebucketObject", &s3.BucketObjectArgs{
			Bucket:               examplebucket.ID(),
			Key:                  pulumi.String("someobject"),
			ServerSideEncryption: pulumi.String("AES256"),
			Source:               pulumi.NewFileAsset("index.html"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

