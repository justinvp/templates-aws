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
		_, err = appautoscaling.NewPolicy(ctx, "ecsPolicy", &appautoscaling.PolicyArgs{
			PolicyType:        pulumi.String("StepScaling"),
			ResourceId:        ecsTarget.ResourceId,
			ScalableDimension: ecsTarget.ScalableDimension,
			ServiceNamespace:  ecsTarget.ServiceNamespace,
			StepScalingPolicyConfiguration: &appautoscaling.PolicyStepScalingPolicyConfigurationArgs{
				AdjustmentType:        pulumi.String("ChangeInCapacity"),
				Cooldown:              pulumi.Int(60),
				MetricAggregationType: pulumi.String("Maximum"),
				StepAdjustment: pulumi.Float64MapArray{
					pulumi.Float64Map{
						"metricIntervalUpperBound": pulumi.Float64(0),
						"scalingAdjustment":        -1,
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

