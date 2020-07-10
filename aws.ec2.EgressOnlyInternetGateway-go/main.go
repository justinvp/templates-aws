package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleVpc, err := ec2.NewVpc(ctx, "exampleVpc", &ec2.VpcArgs{
			AssignGeneratedIpv6CidrBlock: pulumi.Bool(true),
			CidrBlock:                    pulumi.String("10.1.0.0/16"),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewEgressOnlyInternetGateway(ctx, "exampleEgressOnlyInternetGateway", &ec2.EgressOnlyInternetGatewayArgs{
			Tags: pulumi.StringMap{
				"Name": pulumi.String("main"),
			},
			VpcId: exampleVpc.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

