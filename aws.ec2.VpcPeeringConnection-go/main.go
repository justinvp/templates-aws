package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewVpcPeeringConnection(ctx, "foo", &ec2.VpcPeeringConnectionArgs{
			PeerOwnerId: pulumi.Any(_var.Peer_owner_id),
			PeerVpcId:   pulumi.Any(aws_vpc.Bar.Id),
			VpcId:       pulumi.Any(aws_vpc.Foo.Id),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

