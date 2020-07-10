package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.LookupVpcDhcpOptions(ctx, &ec2.LookupVpcDhcpOptionsArgs{
			Filters: []ec2.GetVpcDhcpOptionsFilter{
				ec2.GetVpcDhcpOptionsFilter{
					Name: "key",
					Values: []string{
						"domain-name",
					},
				},
				ec2.GetVpcDhcpOptionsFilter{
					Name: "value",
					Values: []string{
						"example.com",
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

