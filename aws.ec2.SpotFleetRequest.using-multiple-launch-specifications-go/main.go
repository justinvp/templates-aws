package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewSpotFleetRequest(ctx, "foo", &ec2.SpotFleetRequestArgs{
			IamFleetRole: pulumi.String("arn:aws:iam::12345678:role/spot-fleet"),
			LaunchSpecifications: ec2.SpotFleetRequestLaunchSpecificationArray{
				&ec2.SpotFleetRequestLaunchSpecificationArgs{
					Ami:              pulumi.String("ami-d06a90b0"),
					AvailabilityZone: pulumi.String("us-west-2a"),
					InstanceType:     pulumi.String("m1.small"),
					KeyName:          pulumi.String("my-key"),
				},
				&ec2.SpotFleetRequestLaunchSpecificationArgs{
					Ami:              pulumi.String("ami-d06a90b0"),
					AvailabilityZone: pulumi.String("us-west-2a"),
					InstanceType:     pulumi.String("m5.large"),
					KeyName:          pulumi.String("my-key"),
				},
			},
			SpotPrice:      pulumi.String("0.005"),
			TargetCapacity: pulumi.Int(2),
			ValidUntil:     pulumi.String("2019-11-04T20:44:20Z"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

