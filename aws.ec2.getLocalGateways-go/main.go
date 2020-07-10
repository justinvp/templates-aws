package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooLocalGateways, err := ec2.GetLocalGateways(ctx, &ec2.GetLocalGatewaysArgs{
			Tags: map[string]interface{}{
				"service": "production",
			},
		}, nil)
		if err != nil {
			return err
		}
		ctx.Export("foo", fooLocalGateways.Ids)
		return nil
	})
}

