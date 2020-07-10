package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssm.NewParameter(ctx, "foo", &ssm.ParameterArgs{
			Type:  pulumi.String("String"),
			Value: pulumi.String("bar"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

