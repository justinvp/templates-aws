package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := true
		_, err := aws.GetAvailabilityZones(ctx, &aws.GetAvailabilityZonesArgs{
			AllAvailabilityZones: &opt0,
			Filters: []aws.GetAvailabilityZonesFilter{
				aws.GetAvailabilityZonesFilter{
					Name: "opt-in-status",
					Values: []string{
						"not-opted-in",
						"opted-in",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

