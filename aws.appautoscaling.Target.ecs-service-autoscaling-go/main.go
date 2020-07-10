package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/appautoscaling"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appautoscaling.NewTarget(ctx, "ecsTarget", &appautoscaling.TargetArgs{
			MaxCapacity:       pulumi.Int(4),
			MinCapacity:       pulumi.Int(1),
			ResourceId:        pulumi.String(fmt.Sprintf("%v%v%v%v", "service/", aws_ecs_cluster.Example.Name, "/", aws_ecs_service.Example.Name)),
			ScalableDimension: pulumi.String("ecs:service:DesiredCount"),
			ServiceNamespace:  pulumi.String("ecs"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

