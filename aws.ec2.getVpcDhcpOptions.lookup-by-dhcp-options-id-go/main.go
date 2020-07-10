package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := "dopts-12345678"
		_, err := ec2.LookupVpcDhcpOptions(ctx, &ec2.LookupVpcDhcpOptionsArgs{
			DhcpOptionsId: &opt0,
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

