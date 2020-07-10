package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := data.Aws_ec2_local_gateway.Example.Id
		_, err := ec2.GetLocalGatewayVirtualInterfaceGroup(ctx, &ec2.GetLocalGatewayVirtualInterfaceGroupArgs{
			LocalGatewayId: &opt0,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

