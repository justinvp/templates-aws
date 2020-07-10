package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/opsworks"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := opsworks.NewInstance(ctx, "my_instance", &opsworks.InstanceArgs{
			InstanceType: pulumi.String("t2.micro"),
			LayerIds: pulumi.StringArray{
				pulumi.Any(aws_opsworks_custom_layer.My - layer.Id),
			},
			Os:      pulumi.String("Amazon Linux 2015.09"),
			StackId: pulumi.Any(aws_opsworks_stack.Main.Id),
			State:   pulumi.String("stopped"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

