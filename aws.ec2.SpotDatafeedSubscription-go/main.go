package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		defaultBucket, err := s3.NewBucket(ctx, "defaultBucket", nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewSpotDatafeedSubscription(ctx, "defaultSpotDatafeedSubscription", &ec2.SpotDatafeedSubscriptionArgs{
			Bucket: defaultBucket.Bucket,
			Prefix: pulumi.String("my_subdirectory"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

