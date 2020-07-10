package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewDefaultVpc(ctx, "_default", &ec2.DefaultVpcArgs{
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Default VPC"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

