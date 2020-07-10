package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/wafregional"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := wafregional.GetRateBasedMod(ctx, &wafregional.GetRateBasedModArgs{
			Name: "tfWAFRegionalRateBasedRule",
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}

