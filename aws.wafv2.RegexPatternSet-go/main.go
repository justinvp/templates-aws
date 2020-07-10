package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/wafv2"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := wafv2.NewRegexPatternSet(ctx, "example", &wafv2.RegexPatternSetArgs{
			Description: pulumi.String("Example regex pattern set"),
			RegularExpressions: wafv2.RegexPatternSetRegularExpressionArray{
				&wafv2.RegexPatternSetRegularExpressionArgs{
					RegexString: pulumi.String("one"),
				},
				&wafv2.RegexPatternSetRegularExpressionArgs{
					RegexString: pulumi.String("two"),
				},
			},
			Scope: pulumi.String("REGIONAL"),
			Tags: pulumi.StringMap{
				"Tag1": pulumi.String("Value1"),
				"Tag2": pulumi.String("Value2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

