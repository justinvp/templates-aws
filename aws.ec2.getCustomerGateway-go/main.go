package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		foo, err := ec2.LookupCustomerGateway(ctx, &ec2.LookupCustomerGatewayArgs{
			Filters: []ec2.GetCustomerGatewayFilter{
				ec2.GetCustomerGatewayFilter{
					Name: "tag:Name",
					Values: []string{
						"foo-prod",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		main, err := ec2.NewVpnGateway(ctx, "main", &ec2.VpnGatewayArgs{
			AmazonSideAsn: pulumi.String("7224"),
			VpcId:         pulumi.Any(aws_vpc.Main.Id),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewVpnConnection(ctx, "transit", &ec2.VpnConnectionArgs{
			CustomerGatewayId: pulumi.String(foo.Id),
			StaticRoutesOnly:  pulumi.Bool(false),
			Type:              pulumi.String(foo.Type),
			VpnGatewayId:      main.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

