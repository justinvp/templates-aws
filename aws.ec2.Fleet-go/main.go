package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewFleet(ctx, "example", &ec2.FleetArgs{
			LaunchTemplateConfig: &ec2.FleetLaunchTemplateConfigArgs{
				LaunchTemplateSpecification: &ec2.FleetLaunchTemplateConfigLaunchTemplateSpecificationArgs{
					LaunchTemplateId: pulumi.Any(aws_launch_template.Example.Id),
					Version:          pulumi.Any(aws_launch_template.Example.Latest_version),
				},
			},
			TargetCapacitySpecification: &ec2.FleetTargetCapacitySpecificationArgs{
				DefaultTargetCapacityType: pulumi.String("spot"),
				TotalTargetCapacity:       pulumi.Int(5),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

