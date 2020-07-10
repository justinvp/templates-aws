package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/opsworks"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := opsworks.NewNodejsAppLayer(ctx, "app", &opsworks.NodejsAppLayerArgs{
			StackId: pulumi.Any(aws_opsworks_stack.Main.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

