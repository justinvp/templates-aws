package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/opsworks"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := opsworks.NewHaproxyLayer(ctx, "lb", &opsworks.HaproxyLayerArgs{
			StackId:       pulumi.Any(aws_opsworks_stack.Main.Id),
			StatsPassword: pulumi.String("foobarbaz"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

