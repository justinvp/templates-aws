package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/appautoscaling"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appautoscaling.NewTarget(ctx, "replicas", &appautoscaling.TargetArgs{
			MaxCapacity:       pulumi.Int(15),
			MinCapacity:       pulumi.Int(1),
			ResourceId:        pulumi.String(fmt.Sprintf("%v%v", "cluster:", aws_rds_cluster.Example.Id)),
			ScalableDimension: pulumi.String("rds:cluster:ReadReplicaCount"),
			ServiceNamespace:  pulumi.String("rds"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

