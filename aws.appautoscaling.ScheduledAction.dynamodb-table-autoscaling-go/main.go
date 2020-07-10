package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/appautoscaling"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		dynamodbTarget, err := appautoscaling.NewTarget(ctx, "dynamodbTarget", &appautoscaling.TargetArgs{
			MaxCapacity:       pulumi.Int(100),
			MinCapacity:       pulumi.Int(5),
			ResourceId:        pulumi.String("table/tableName"),
			ScalableDimension: pulumi.String("dynamodb:table:ReadCapacityUnits"),
			ServiceNamespace:  pulumi.String("dynamodb"),
		})
		if err != nil {
			return err
		}
		_, err = appautoscaling.NewScheduledAction(ctx, "dynamodbScheduledAction", &appautoscaling.ScheduledActionArgs{
			ResourceId:        dynamodbTarget.ResourceId,
			ScalableDimension: dynamodbTarget.ScalableDimension,
			ScalableTargetAction: &appautoscaling.ScheduledActionScalableTargetActionArgs{
				MaxCapacity: pulumi.Int(200),
				MinCapacity: pulumi.Int(1),
			},
			Schedule:         pulumi.String("at(2006-01-02T15:04:05)"),
			ServiceNamespace: dynamodbTarget.ServiceNamespace,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

