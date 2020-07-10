package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudfront"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		selected, err := s3.LookupBucket(ctx, &s3.LookupBucketArgs{
			Bucket: "a-test-bucket",
		}, nil)
		if err != nil {
			return err
		}
		_, err = cloudfront.NewDistribution(ctx, "test", &cloudfront.DistributionArgs{
			Origins: cloudfront.DistributionOriginArray{
				&cloudfront.DistributionOriginArgs{
					DomainName: pulumi.String(selected.BucketDomainName),
					OriginId:   pulumi.String("s3-selected-bucket"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

