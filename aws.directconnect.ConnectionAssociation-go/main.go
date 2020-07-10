package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/directconnect"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleConnection, err := directconnect.NewConnection(ctx, "exampleConnection", &directconnect.ConnectionArgs{
			Bandwidth: pulumi.String("1Gbps"),
			Location:  pulumi.String("EqSe2"),
		})
		if err != nil {
			return err
		}
		exampleLinkAggregationGroup, err := directconnect.NewLinkAggregationGroup(ctx, "exampleLinkAggregationGroup", &directconnect.LinkAggregationGroupArgs{
			ConnectionsBandwidth: pulumi.String("1Gbps"),
			Location:             pulumi.String("EqSe2"),
		})
		if err != nil {
			return err
		}
		_, err = directconnect.NewConnectionAssociation(ctx, "exampleConnectionAssociation", &directconnect.ConnectionAssociationArgs{
			ConnectionId: exampleConnection.ID(),
			LagId:        exampleLinkAggregationGroup.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

