package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.NewParameterGroup(ctx, "_default", &rds.ParameterGroupArgs{
			Family: pulumi.String("mysql5.6"),
			Parameters: rds.ParameterGroupParameterArray{
				&rds.ParameterGroupParameterArgs{
					Name:  pulumi.String("character_set_server"),
					Value: pulumi.String("utf8"),
				},
				&rds.ParameterGroupParameterArgs{
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

