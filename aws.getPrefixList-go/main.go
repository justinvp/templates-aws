package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		privateS3VpcEndpoint, err := ec2.NewVpcEndpoint(ctx, "privateS3VpcEndpoint", &ec2.VpcEndpointArgs{
			ServiceName: pulumi.String("com.amazonaws.us-west-2.s3"),
			VpcId:       pulumi.Any(aws_vpc.Foo.Id),
		})
		if err != nil {
			return err
		}
		bar, err := ec2.NewNetworkAcl(ctx, "bar", &ec2.NetworkAclArgs{
			VpcId: pulumi.Any(aws_vpc.Foo.Id),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewNetworkAclRule(ctx, "privateS3NetworkAclRule", &ec2.NetworkAclRuleArgs{
			CidrBlock: privateS3PrefixList.ApplyT(func(privateS3PrefixList aws.GetPrefixListResult) (string, error) {
				return privateS3PrefixList.CidrBlocks[0], nil
			}).(pulumi.StringOutput),
			Egress:       pulumi.Bool(false),
			FromPort:     pulumi.Int(443),
			NetworkAclId: bar.ID(),
			Protocol:     pulumi.String("tcp"),
			RuleAction:   pulumi.String("allow"),
			RuleNumber:   pulumi.Int(200),
			ToPort:       pulumi.Int(443),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

