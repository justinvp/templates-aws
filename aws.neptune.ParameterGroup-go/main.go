package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/neptune"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := neptune.NewParameterGroup(ctx, "example", &neptune.ParameterGroupArgs{
			Family: pulumi.String("neptune1"),
			Parameters: neptune.ParameterGroupParameterArray{
				&neptune.ParameterGroupParameterArgs{
					Name:  pulumi.String("neptune_query_timeout"),
					Value: pulumi.String("25"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

