package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := "availability-zone-id"
		_, err := ec2.GetInstanceTypeOfferings(ctx, &ec2.GetInstanceTypeOfferingsArgs{
			Filters: []ec2.GetInstanceTypeOfferingsFilter{
				ec2.GetInstanceTypeOfferingsFilter{
					Name: "instance-type",
					Values: []string{
						"t2.micro",
						"t3.micro",
					},
				},
				ec2.GetInstanceTypeOfferingsFilter{
					Name: "location",
					Values: []string{
						"usw2-az4",
					},
				},
			},
			LocationType: &opt0,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

