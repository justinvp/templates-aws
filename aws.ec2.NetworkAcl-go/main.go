package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewNetworkAcl(ctx, "main", &ec2.NetworkAclArgs{
			Egress: ec2.NetworkAclEgressArray{
				&ec2.NetworkAclEgressArgs{
					Action:    pulumi.String("allow"),
					CidrBlock: pulumi.String("10.3.0.0/18"),
					FromPort:  pulumi.Int(443),
					Protocol:  pulumi.String("tcp"),
					RuleNo:    pulumi.Int(200),
					ToPort:    pulumi.Int(443),
				},
			},
			Ingress: ec2.NetworkAclIngressArray{
				&ec2.NetworkAclIngressArgs{
					Action:    pulumi.String("allow"),
					CidrBlock: pulumi.String("10.3.0.0/18"),
					FromPort:  pulumi.Int(80),
					Protocol:  pulumi.String("tcp"),
					RuleNo:    pulumi.Int(100),
					ToPort:    pulumi.Int(80),
				},
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("main"),
			},
			VpcId: pulumi.Any(aws_vpc.Main.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

