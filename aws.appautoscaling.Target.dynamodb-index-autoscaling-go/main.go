package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/appautoscaling"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appautoscaling.NewTarget(ctx, "dynamodbIndexReadTarget", &appautoscaling.TargetArgs{
			MaxCapacity:       pulumi.Int(100),
			MinCapacity:       pulumi.Int(5),
			ResourceId:        pulumi.String(fmt.Sprintf("%v%v%v%v", "table/", aws_dynamodb_table.Example.Name, "/index/", _var.Index_name)),
			ScalableDimension: pulumi.String("dynamodb:index:ReadCapacityUnits"),
			ServiceNamespace:  pulumi.String("dynamodb"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

