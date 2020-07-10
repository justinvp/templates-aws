package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.NewInstance(ctx, "example", &rds.InstanceArgs{
			AllocatedStorage:    pulumi.Int(50),
			MaxAllocatedStorage: pulumi.Int(100),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

