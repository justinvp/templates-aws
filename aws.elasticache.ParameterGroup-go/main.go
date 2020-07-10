package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elasticache"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elasticache.NewParameterGroup(ctx, "_default", &elasticache.ParameterGroupArgs{
			Family: pulumi.String("redis2.8"),
			Parameters: elasticache.ParameterGroupParameterArray{
				&elasticache.ParameterGroupParameterArgs{
					Name:  pulumi.String("activerehashing"),
					Value: pulumi.String("yes"),
				},
				&elasticache.ParameterGroupParameterArgs{
					Name:  pulumi.String("min-slaves-to-write"),
					Value: pulumi.String("2"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

