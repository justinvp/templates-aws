package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/dax"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dax.NewParameterGroup(ctx, "example", &dax.ParameterGroupArgs{
			Parameters: dax.ParameterGroupParameterArray{
				&dax.ParameterGroupParameterArgs{
					Name:  pulumi.String("query-ttl-millis"),
					Value: pulumi.String("100000"),
				},
				&dax.ParameterGroupParameterArgs{
					Name:  pulumi.String("record-ttl-millis"),
					Value: pulumi.String("100000"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

