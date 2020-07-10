package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewSecurityGroupRule(ctx, "example", &ec2.SecurityGroupRuleArgs{
			Type:            pulumi.String("ingress"),
			FromPort:        pulumi.Int(0),
			ToPort:          pulumi.Int(65535),
			Protocol:        pulumi.String("tcp"),
			CidrBlocks:      aws_vpc.Example.Cidr_block,
			SecurityGroupId: pulumi.String("sg-123456"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

