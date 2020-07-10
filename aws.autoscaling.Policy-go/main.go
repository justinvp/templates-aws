package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/autoscaling"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		bar, err := autoscaling.NewGroup(ctx, "bar", &autoscaling.GroupArgs{
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-east-1a"),
			},
			ForceDelete:            pulumi.Bool(true),
			HealthCheckGracePeriod: pulumi.Int(300),
			HealthCheckType:        pulumi.String("ELB"),
			LaunchConfiguration:    pulumi.Any(aws_launch_configuration.Foo.Name),
			MaxSize:                pulumi.Int(5),
			MinSize:                pulumi.Int(2),
		})
		if err != nil {
			return err
		}
		_, err = autoscaling.NewPolicy(ctx, "bat", &autoscaling.PolicyArgs{
			AdjustmentType:       pulumi.String("ChangeInCapacity"),
			AutoscalingGroupName: bar.Name,
			Cooldown:             pulumi.Int(300),
			ScalingAdjustment:    pulumi.Int(4),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

