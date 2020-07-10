package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewVpcEndpointSubnetAssociation(ctx, "snEc2", &ec2.VpcEndpointSubnetAssociationArgs{
			SubnetId:      pulumi.Any(aws_subnet.Sn.Id),
			VpcEndpointId: pulumi.Any(aws_vpc_endpoint.Ec2.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

