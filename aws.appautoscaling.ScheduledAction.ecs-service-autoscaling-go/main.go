package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/appautoscaling"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		ecsTarget, err := appautoscaling.NewTarget(ctx, "ecsTarget", &appautoscaling.TargetArgs{
			MaxCapacity:       pulumi.Int(4),
			MinCapacity:       pulumi.Int(1),
			ResourceId:        pulumi.String("service/clusterName/serviceName"),
			ScalableDimension: pulumi.String("ecs:service:DesiredCount"),
			ServiceNamespace:  pulumi.String("ecs"),
		})
		if err != nil {
			return err
		}
		_, err = appautoscaling.NewScheduledAction(ctx, "ecsScheduledAction", &appautoscaling.ScheduledActionArgs{
			ResourceId:        ecsTarget.ResourceId,
			ScalableDimension: ecsTarget.ScalableDimension,
			ScalableTargetAction: &appautoscaling.ScheduledActionScalableTargetActionArgs{
				MaxCapacity: pulumi.Int(10),
				MinCapacity: pulumi.Int(1),
			},
			Schedule:         pulumi.String("at(2006-01-02T15:04:05)"),
			ServiceNamespace: ecsTarget.ServiceNamespace,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

