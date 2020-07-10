package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/servicediscovery"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleVpc, err := ec2.NewVpc(ctx, "exampleVpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		})
		if err != nil {
			return err
		}
		_, err = servicediscovery.NewPrivateDnsNamespace(ctx, "examplePrivateDnsNamespace", &servicediscovery.PrivateDnsNamespaceArgs{
			Description: pulumi.String("example"),
			Vpc:         exampleVpc.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

