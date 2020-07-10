package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewSpotFleetRequest(ctx, "cheapCompute", &ec2.SpotFleetRequestArgs{
			AllocationStrategy: pulumi.String("diversified"),
			IamFleetRole:       pulumi.String("arn:aws:iam::12345678:role/spot-fleet"),
			LaunchSpecifications: ec2.SpotFleetRequestLaunchSpecificationArray{
				&ec2.SpotFleetRequestLaunchSpecificationArgs{
					Ami:                   pulumi.String("ami-1234"),
					IamInstanceProfileArn: pulumi.Any(aws_iam_instance_profile.Example.Arn),
					InstanceType:          pulumi.String("m4.10xlarge"),
					PlacementTenancy:      pulumi.String("dedicated"),
					SpotPrice:             pulumi.String("2.793"),
				},
				&ec2.SpotFleetRequestLaunchSpecificationArgs{
					Ami:                   pulumi.String("ami-5678"),
					AvailabilityZone:      pulumi.String("us-west-1a"),
					IamInstanceProfileArn: pulumi.Any(aws_iam_instance_profile.Example.Arn),
					InstanceType:          pulumi.String("m4.4xlarge"),
					KeyName:               pulumi.String("my-key"),
					RootBlockDevice: pulumi.StringMapArray{
						pulumi.StringMap{
							"volumeSize": pulumi.String("300"),
							"volumeType": pulumi.String("gp2"),
						},
					},
					SpotPrice: pulumi.String("1.117"),
					SubnetId:  pulumi.String("subnet-1234"),
					Tags: pulumi.StringMap{
						"Name": pulumi.String("spot-fleet-example"),
					},
					WeightedCapacity: pulumi.String("35"),
				},
			},
			SpotPrice:      pulumi.String("0.03"),
			TargetCapacity: pulumi.Int(6),
			ValidUntil:     pulumi.String("2019-11-04T20:44:20Z"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

