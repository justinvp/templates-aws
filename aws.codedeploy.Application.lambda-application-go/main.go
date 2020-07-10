package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/codedeploy"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := codedeploy.NewApplication(ctx, "example", &codedeploy.ApplicationArgs{
			ComputePlatform: pulumi.String("Lambda"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

