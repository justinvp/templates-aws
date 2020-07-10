package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewMainRouteTableAssociation(ctx, "mainRouteTableAssociation", &ec2.MainRouteTableAssociationArgs{
			RouteTableId: pulumi.Any(aws_route_table.Bar.Id),
			VpcId:        pulumi.Any(aws_vpc.Foo.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

