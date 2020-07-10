package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.NewClusterParameterGroup(ctx, "_default", &rds.ClusterParameterGroupArgs{
			Description: pulumi.String("RDS default cluster parameter group"),
			Family:      pulumi.String("aurora5.6"),
			Parameters: rds.ClusterParameterGroupParameterArray{
				&rds.ClusterParameterGroupParameterArgs{
					Name:  pulumi.String("character_set_server"),
					Value: pulumi.String("utf8"),
				},
				&rds.ClusterParameterGroupParameterArgs{
					Name:  pulumi.String("character_set_client"),
					Value: pulumi.String("utf8"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

