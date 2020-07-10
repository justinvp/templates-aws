package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/waf"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleRegexPatternSet, err := waf.NewRegexPatternSet(ctx, "exampleRegexPatternSet", &waf.RegexPatternSetArgs{
			RegexPatternStrings: pulumi.StringArray{
				pulumi.String("one"),
				pulumi.String("two"),
			},
		})
		if err != nil {
			return err
		}
		_, err = waf.NewRegexMatchSet(ctx, "exampleRegexMatchSet", &waf.RegexMatchSetArgs{
			RegexMatchTuples: waf.RegexMatchSetRegexMatchTupleArray{
				&waf.RegexMatchSetRegexMatchTupleArgs{
					FieldToMatch: &waf.RegexMatchSetRegexMatchTupleFieldToMatchArgs{
						Data: pulumi.String("User-Agent"),
						Type: pulumi.String("HEADER"),
					},
					RegexPatternSetId:  exampleRegexPatternSet.ID(),
					TextTransformation: pulumi.String("NONE"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

