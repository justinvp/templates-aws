package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2transitgateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := aws_dx_gateway.Example.Id
		opt1 := aws_ec2_transit_gateway.Example.Id
		_, err := ec2transitgateway.GetDirectConnectGatewayAttachment(ctx, &ec2transitgateway.GetDirectConnectGatewayAttachmentArgs{
			DxGatewayId:      &opt0,
			TransitGatewayId: &opt1,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

