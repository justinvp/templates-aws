package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := _var.Vpc_id
		exampleNetworkAcls, err := ec2.GetNetworkAcls(ctx, &ec2.GetNetworkAclsArgs{
			VpcId: &opt0,
		}, nil)
		if err != nil {
			return err
		}
		ctx.Export("example", exampleNetworkAcls.Ids)
		return nil
	})
}

