package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cfg"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cfg.NewConfigurationAggregator(ctx, "account", &cfg.ConfigurationAggregatorArgs{
			AccountAggregationSource: &cfg.ConfigurationAggregatorAccountAggregationSourceArgs{
				AccountIds: pulumi.StringArray{
					pulumi.String("123456789012"),
				},
				Regions: pulumi.StringArray{
					pulumi.String("us-west-2"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

