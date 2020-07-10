package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/autoscaling"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		foobarGroup, err := autoscaling.NewGroup(ctx, "foobarGroup", &autoscaling.GroupArgs{
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-west-2a"),
			},
			ForceDelete:            pulumi.Bool(true),
			HealthCheckGracePeriod: pulumi.Int(300),
			HealthCheckType:        pulumi.String("ELB"),
			MaxSize:                pulumi.Int(1),
			MinSize:                pulumi.Int(1),
			TerminationPolicies: pulumi.StringArray{
				pulumi.String("OldestInstance"),
			},
		})
		if err != nil {
			return err
		}
		_, err = autoscaling.NewSchedule(ctx, "foobarSchedule", &autoscaling.ScheduleArgs{
			AutoscalingGroupName: foobarGroup.Name,
			DesiredCapacity:      pulumi.Int(0),
			EndTime:              pulumi.String("2016-12-12T06:00:00Z"),
			MaxSize:              pulumi.Int(1),
			MinSize:              pulumi.Int(0),
			ScheduledActionName:  pulumi.String("foobar"),
			StartTime:            pulumi.String("2016-12-11T18:00:00Z"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

