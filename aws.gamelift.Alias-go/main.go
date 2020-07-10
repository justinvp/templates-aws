package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/gamelift"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := gamelift.NewAlias(ctx, "example", &gamelift.AliasArgs{
			Description: pulumi.String("Example Description"),
			RoutingStrategy: &gamelift.AliasRoutingStrategyArgs{
				Message: pulumi.String("Example Message"),
				Type:    pulumi.String("TERMINAL"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

