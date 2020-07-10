package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/xray"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := xray.NewSamplingRule(ctx, "example", &xray.SamplingRuleArgs{
			Attributes: pulumi.StringMap{
				"Hello": pulumi.String("Tris"),
			},
			FixedRate:     pulumi.Float64(0.05),
			Host:          pulumi.String("*"),
			HttpMethod:    pulumi.String("*"),
			Priority:      pulumi.Int(10000),
			ReservoirSize: pulumi.Int(1),
			ResourceArn:   pulumi.String("*"),
			RuleName:      pulumi.String("example"),
			ServiceName:   pulumi.String("*"),
			ServiceType:   pulumi.String("*"),
			UrlPath:       pulumi.String("*"),
			Version:       pulumi.Int(1),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

