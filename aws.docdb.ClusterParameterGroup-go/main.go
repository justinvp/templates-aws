package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/docdb"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := docdb.NewClusterParameterGroup(ctx, "example", &docdb.ClusterParameterGroupArgs{
			Description: pulumi.String("docdb cluster parameter group"),
			Family:      pulumi.String("docdb3.6"),
			Parameters: docdb.ClusterParameterGroupParameterArray{
				&docdb.ClusterParameterGroupParameterArgs{
					Name:  pulumi.String("tls"),
					Value: pulumi.String("enabled"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

