package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewRouteTableAssociation(ctx, "routeTableAssociation", &ec2.RouteTableAssociationArgs{
			SubnetId:     pulumi.Any(aws_subnet.Foo.Id),
			RouteTableId: pulumi.Any(aws_route_table.Bar.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

