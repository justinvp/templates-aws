package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		secondaryCidr, err := ec2.NewVpcIpv4CidrBlockAssociation(ctx, "secondaryCidr", &ec2.VpcIpv4CidrBlockAssociationArgs{
			CidrBlock: pulumi.String("172.2.0.0/16"),
			VpcId:     pulumi.Any(aws_vpc.Main.Id),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewSubnet(ctx, "inSecondaryCidr", &ec2.SubnetArgs{
			CidrBlock: pulumi.String("172.2.0.0/24"),
			VpcId:     secondaryCidr.VpcId,
		})
		if err != nil {
			return err
		}
		return nil
	})
}

