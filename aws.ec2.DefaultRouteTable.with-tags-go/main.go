package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewDefaultRouteTable(ctx, "defaultRouteTable", &ec2.DefaultRouteTableArgs{
			DefaultRouteTableId: pulumi.Any(aws_vpc.Foo.Default_route_table_id),
			Routes: ec2.DefaultRouteTableRouteArray{
				nil,
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("default table"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

