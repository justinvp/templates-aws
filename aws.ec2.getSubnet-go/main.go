package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := subnetId
		selected, err := ec2.LookupSubnet(ctx, &ec2.LookupSubnetArgs{
			Id: &opt0,
		}, nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewSecurityGroup(ctx, "subnet", &ec2.SecurityGroupArgs{
			Ingress: ec2.SecurityGroupIngressArray{
				&ec2.SecurityGroupIngressArgs{
					CidrBlocks: pulumi.StringArray{
						pulumi.String(selected.CidrBlock),
					},
					FromPort: pulumi.Int(80),
					Protocol: pulumi.String("tcp"),
					ToPort:   pulumi.Int(80),
				},
			},
			VpcId: pulumi.String(selected.VpcId),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

