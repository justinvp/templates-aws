package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cfg"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cfg.NewAggregateAuthorization(ctx, "example", &cfg.AggregateAuthorizationArgs{
			AccountId: pulumi.String("123456789012"),
			Region:    pulumi.String("eu-west-2"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

