package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2transitgateway"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/providers"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := providers.Newaws(ctx, "local", &providers.awsArgs{
			Region: pulumi.String("us-east-1"),
		})
		if err != nil {
			return err
		}
		_, err = providers.Newaws(ctx, "peer", &providers.awsArgs{
			Region: pulumi.String("us-west-2"),
		})
		if err != nil {
			return err
		}
		peerRegion, err := aws.GetRegion(ctx, nil, nil)
		if err != nil {
			return err
		}
		localTransitGateway, err := ec2transitgateway.NewTransitGateway(ctx, "localTransitGateway", &ec2transitgateway.TransitGatewayArgs{
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Local TGW"),
			},
		}, pulumi.Provider(aws.Local))
		if err != nil {
			return err
		}
		peerTransitGateway, err := ec2transitgateway.NewTransitGateway(ctx, "peerTransitGateway", &ec2transitgateway.TransitGatewayArgs{
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Peer TGW"),
			},
		}, pulumi.Provider(aws.Peer))
		if err != nil {
			return err
		}
		_, err = ec2transitgateway.NewPeeringAttachment(ctx, "example", &ec2transitgateway.PeeringAttachmentArgs{
			PeerAccountId:        peerTransitGateway.OwnerId,
			PeerRegion:           pulumi.String(peerRegion.Name),
			PeerTransitGatewayId: peerTransitGateway.ID(),
			TransitGatewayId:     localTransitGateway.ID(),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("TGW Peering Requestor"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

