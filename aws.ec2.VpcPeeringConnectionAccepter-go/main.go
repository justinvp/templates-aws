package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/providers"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := providers.Newaws(ctx, "peer", &providers.awsArgs{
			Region: pulumi.String("us-west-2"),
		})
		if err != nil {
			return err
		}
		main, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		})
		if err != nil {
			return err
		}
		peerVpc, err := ec2.NewVpc(ctx, "peerVpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.1.0.0/16"),
		}, pulumi.Provider("aws.peer"))
		if err != nil {
			return err
		}
		peerCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		peerVpcPeeringConnection, err := ec2.NewVpcPeeringConnection(ctx, "peerVpcPeeringConnection", &ec2.VpcPeeringConnectionArgs{
			AutoAccept:  pulumi.Bool(false),
			PeerOwnerId: pulumi.String(peerCallerIdentity.AccountId),
			PeerRegion:  pulumi.String("us-west-2"),
			PeerVpcId:   peerVpc.ID(),
			Tags: pulumi.StringMap{
				"Side": pulumi.String("Requester"),
			},
			VpcId: main.ID(),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewVpcPeeringConnectionAccepter(ctx, "peerVpcPeeringConnectionAccepter", &ec2.VpcPeeringConnectionAccepterArgs{
			AutoAccept: pulumi.Bool(true),
			Tags: pulumi.StringMap{
				"Side": pulumi.String("Accepter"),
			},
			VpcPeeringConnectionId: peerVpcPeeringConnection.ID(),
		}, pulumi.Provider("aws.peer"))
		if err != nil {
			return err
		}
		return nil
	})
}

