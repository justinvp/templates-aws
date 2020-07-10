package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleNetworkInterfaces, err := ec2.GetNetworkInterfaces(ctx, nil, nil)
		if err != nil {
			return err
		}
		ctx.Export("example", exampleNetworkInterfaces.Ids)
		return nil
	})
}

