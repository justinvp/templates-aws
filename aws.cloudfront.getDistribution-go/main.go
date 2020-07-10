package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloudfront"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudfront.LookupDistribution(ctx, &cloudfront.LookupDistributionArgs{
			Id: "EDFDVBD632BHDS5",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

