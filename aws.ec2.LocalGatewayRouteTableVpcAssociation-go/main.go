package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := "arn:aws:outposts:us-west-2:123456789012:outpost/op-1234567890abcdef"
		exampleLocalGatewayRouteTable, err := ec2.GetLocalGatewayRouteTable(ctx, &ec2.GetLocalGatewayRouteTableArgs{
			OutpostArn: &opt0,
		}, nil)
		if err != nil {
			return err
		}
		exampleVpc, err := ec2.NewVpc(ctx, "exampleVpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewLocalGatewayRouteTableVpcAssociation(ctx, "exampleLocalGatewayRouteTableVpcAssociation", &ec2.LocalGatewayRouteTableVpcAssociationArgs{
			LocalGatewayRouteTableId: pulumi.String(exampleLocalGatewayRouteTable.Id),
			VpcId:                    exampleVpc.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

