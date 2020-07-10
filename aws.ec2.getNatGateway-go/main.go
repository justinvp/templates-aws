package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := aws_subnet.Public.Id
		_, err := ec2.LookupNatGateway(ctx, &ec2.LookupNatGatewayArgs{
			SubnetId: &opt0,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

