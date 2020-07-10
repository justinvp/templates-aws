package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewVpcEndpointRouteTableAssociation(ctx, "example", &ec2.VpcEndpointRouteTableAssociationArgs{
			RouteTableId:  pulumi.Any(aws_route_table.Example.Id),
			VpcEndpointId: pulumi.Any(aws_vpc_endpoint.Example.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

