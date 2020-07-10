package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewSubnet(ctx, "main", &ec2.SubnetArgs{
			CidrBlock: pulumi.String("10.0.1.0/24"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Main"),
			},
			VpcId: pulumi.Any(aws_vpc.Main.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

