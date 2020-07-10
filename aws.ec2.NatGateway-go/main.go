package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewNatGateway(ctx, "gw", &ec2.NatGatewayArgs{
			AllocationId: pulumi.Any(aws_eip.Nat.Id),
			SubnetId:     pulumi.Any(aws_subnet.Example.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

