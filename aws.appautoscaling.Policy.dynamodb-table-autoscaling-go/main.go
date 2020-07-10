package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/appautoscaling"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		dynamodbTableReadTarget, err := appautoscaling.NewTarget(ctx, "dynamodbTableReadTarget", &appautoscaling.TargetArgs{
			MaxCapacity:       pulumi.Int(100),
			MinCapacity:       pulumi.Int(5),
			ResourceId:        pulumi.String("table/tableName"),
			ScalableDimension: pulumi.String("dynamodb:table:ReadCapacityUnits"),
			ServiceNamespace:  pulumi.String("dynamodb"),
		})
		if err != nil {
			return err
		}
		_, err = appautoscaling.NewPolicy(ctx, "dynamodbTableReadPolicy", &appautoscaling.PolicyArgs{
			PolicyType:        pulumi.String("TargetTrackingScaling"),
			ResourceId:        dynamodbTableReadTarget.ResourceId,
			ScalableDimension: dynamodbTableReadTarget.ScalableDimension,
			ServiceNamespace:  dynamodbTableReadTarget.ServiceNamespace,
			TargetTrackingScalingPolicyConfiguration: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationArgs{
				PredefinedMetricSpecification: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationArgs{
					PredefinedMetricType: pulumi.String("DynamoDBReadCapacityUtilization"),
				},
				TargetValue: pulumi.Float64(70),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

