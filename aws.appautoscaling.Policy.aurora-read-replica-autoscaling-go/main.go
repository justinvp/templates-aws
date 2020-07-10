package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/appautoscaling"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		replicasTarget, err := appautoscaling.NewTarget(ctx, "replicasTarget", &appautoscaling.TargetArgs{
			MaxCapacity:       pulumi.Int(15),
			MinCapacity:       pulumi.Int(1),
			ResourceId:        pulumi.String(fmt.Sprintf("%v%v", "cluster:", aws_rds_cluster.Example.Id)),
			ScalableDimension: pulumi.String("rds:cluster:ReadReplicaCount"),
			ServiceNamespace:  pulumi.String("rds"),
		})
		if err != nil {
			return err
		}
		_, err = appautoscaling.NewPolicy(ctx, "replicasPolicy", &appautoscaling.PolicyArgs{
			PolicyType:        pulumi.String("TargetTrackingScaling"),
			ResourceId:        replicasTarget.ResourceId,
			ScalableDimension: replicasTarget.ScalableDimension,
			ServiceNamespace:  replicasTarget.ServiceNamespace,
			TargetTrackingScalingPolicyConfiguration: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationArgs{
				PredefinedMetricSpecification: &appautoscaling.PolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecificationArgs{
					PredefinedMetricType: pulumi.String("RDSReaderAverageCPUUtilization"),
				},
				ScaleInCooldown:  pulumi.Int(300),
				ScaleOutCooldown: pulumi.Int(300),
				TargetValue:      pulumi.Float64(75),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

