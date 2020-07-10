package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		barNetworkAcl, err := ec2.NewNetworkAcl(ctx, "barNetworkAcl", &ec2.NetworkAclArgs{
			VpcId: pulumi.Any(aws_vpc.Foo.Id),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewNetworkAclRule(ctx, "barNetworkAclRule", &ec2.NetworkAclRuleArgs{
			NetworkAclId: barNetworkAcl.ID(),
			RuleNumber:   pulumi.Int(200),
			Egress:       pulumi.Bool(false),
			Protocol:     pulumi.String("tcp"),
			RuleAction:   pulumi.String("allow"),
			CidrBlock:    pulumi.Any(aws_vpc.Foo.Cidr_block),
			FromPort:     pulumi.Int(22),
			ToPort:       pulumi.Int(22),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

