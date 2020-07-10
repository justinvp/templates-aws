package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		opt0 := "10.0.1.0/22"
		opt1 := aws_vpc.Foo.Id
		pc, err := ec2.LookupVpcPeeringConnection(ctx, &ec2.LookupVpcPeeringConnectionArgs{
			PeerCidrBlock: &opt0,
			VpcId:         &opt1,
		}, nil)
		if err != nil {
			return err
		}
		rt, err := ec2.NewRouteTable(ctx, "rt", &ec2.RouteTableArgs{
			VpcId: pulumi.Any(aws_vpc.Foo.Id),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewRoute(ctx, "route", &ec2.RouteArgs{
			DestinationCidrBlock:   pulumi.String(pc.PeerCidrBlock),
			RouteTableId:           rt.ID(),
			VpcPeeringConnectionId: pulumi.String(pc.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

