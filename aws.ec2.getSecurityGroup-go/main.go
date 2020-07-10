package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := securityGroupId
		selected, err := ec2.LookupSecurityGroup(ctx, &ec2.LookupSecurityGroupArgs{
			Id: &opt0,
		}, nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewSubnet(ctx, "subnet", &ec2.SubnetArgs{
			CidrBlock: pulumi.String("10.0.1.0/24"),
			VpcId:     pulumi.String(selected.VpcId),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

