package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/waf"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := waf.NewRegexPatternSet(ctx, "example", &waf.RegexPatternSetArgs{
			RegexPatternStrings: pulumi.StringArray{
				pulumi.String("one"),
				pulumi.String("two"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

