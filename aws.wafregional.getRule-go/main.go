package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/wafregional"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := wafregional.LookupRule(ctx, &wafregional.LookupRuleArgs{
			Name: "tfWAFRegionalRule",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

