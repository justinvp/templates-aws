package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewPlacementGroup(ctx, "web", &ec2.PlacementGroupArgs{
			Strategy: pulumi.String("cluster"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

