package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewRouteTable(ctx, "routeTable", &ec2.RouteTableArgs{
			Routes: ec2.RouteTableRouteArray{
				&ec2.RouteTableRouteArgs{
					CidrBlock: pulumi.String("10.0.1.0/24"),
					GatewayId: pulumi.Any(aws_internet_gateway.Main.Id),
				},
				&ec2.RouteTableRouteArgs{
					EgressOnlyGatewayId: pulumi.Any(aws_egress_only_internet_gateway.Foo.Id),
					Ipv6CidrBlock:       pulumi.String("::/0"),
				},
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("main"),
			},
			VpcId: pulumi.Any(aws_vpc.Default.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

