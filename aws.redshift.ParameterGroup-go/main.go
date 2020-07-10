package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/redshift"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := redshift.NewParameterGroup(ctx, "bar", &redshift.ParameterGroupArgs{
			Family: pulumi.String("redshift-1.0"),
			Parameters: redshift.ParameterGroupParameterArray{
				&redshift.ParameterGroupParameterArgs{
					Name:  pulumi.String("require_ssl"),
					Value: pulumi.String("true"),
				},
				&redshift.ParameterGroupParameterArgs{
					Name:  pulumi.String("query_group"),
					Value: pulumi.String("example"),
				},
				&redshift.ParameterGroupParameterArgs{
					Name:  pulumi.String("enable_user_activity_logging"),
					Value: pulumi.String("true"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

