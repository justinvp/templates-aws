package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/waf"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := waf.LookupRateBasedRule(ctx, &waf.LookupRateBasedRuleArgs{
			Name: "tfWAFRateBasedRule",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

