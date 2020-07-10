package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewCustomerGateway(ctx, "main", &ec2.CustomerGatewayArgs{
			BgpAsn:    pulumi.Int(65000),
			IpAddress: pulumi.String("172.83.124.10"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("main-customer-gateway"),
			},
			Type: pulumi.String("ipsec.1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

