package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := awsEc2LocalGatewayRouteTable
		_, err := ec2.GetLocalGatewayRouteTable(ctx, &ec2.GetLocalGatewayRouteTableArgs{
			LocalGatewayRouteTableId: &opt0,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

