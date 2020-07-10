package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cloud9"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloud9.NewEnvironmentEC2(ctx, "example", &cloud9.EnvironmentEC2Args{
			InstanceType: pulumi.String("t2.micro"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

