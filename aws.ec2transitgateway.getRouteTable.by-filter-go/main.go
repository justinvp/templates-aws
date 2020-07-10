package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2transitgateway"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2transitgateway.LookupRouteTable(ctx, &ec2transitgateway.LookupRouteTableArgs{
			Filters: []ec2transitgateway.GetRouteTableFilter{
				ec2transitgateway.GetRouteTableFilter{
					Name: "default-association-route-table",
					Values: []string{
						"true",
					},
				},
				ec2transitgateway.GetRouteTableFilter{
					Name: "transit-gateway-id",
					Values: []string{
						"tgw-12345678",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

