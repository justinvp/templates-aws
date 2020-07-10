package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elasticbeanstalk"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		defaultBucket, err := s3.NewBucket(ctx, "defaultBucket", nil)
		if err != nil {
			return err
		}
		defaultBucketObject, err := s3.NewBucketObject(ctx, "defaultBucketObject", &s3.BucketObjectArgs{
			Bucket: defaultBucket.ID(),
			Key:    pulumi.String("beanstalk/go-v1.zip"),
			Source: pulumi.NewFileAsset("go-v1.zip"),
		})
		if err != nil {
			return err
		}
		_, err = elasticbeanstalk.NewApplication(ctx, "defaultApplication", &elasticbeanstalk.ApplicationArgs{
			Description: pulumi.String("tf-test-desc"),
		})
		if err != nil {
			return err
		}
		_, err = elasticbeanstalk.NewApplicationVersion(ctx, "defaultApplicationVersion", &elasticbeanstalk.ApplicationVersionArgs{
			Application: pulumi.String("tf-test-name"),
			Bucket:      defaultBucket.ID(),
			Description: pulumi.String("application version"),
			Key:         defaultBucketObject.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

