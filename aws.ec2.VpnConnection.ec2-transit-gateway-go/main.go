package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2transitgateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleTransitGateway, err := ec2transitgateway.NewTransitGateway(ctx, "exampleTransitGateway", nil)
		if err != nil {
			return err
		}
		exampleCustomerGateway, err := ec2.NewCustomerGateway(ctx, "exampleCustomerGateway", &ec2.CustomerGatewayArgs{
			BgpAsn:    pulumi.Int(65000),
			IpAddress: pulumi.String("172.0.0.1"),
			Type:      pulumi.String("ipsec.1"),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewVpnConnection(ctx, "exampleVpnConnection", &ec2.VpnConnectionArgs{
			CustomerGatewayId: exampleCustomerGateway.ID(),
			TransitGatewayId:  exampleTransitGateway.ID(),
			Type:              exampleCustomerGateway.Type,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

