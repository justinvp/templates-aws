package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/wafregional"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleRegexPatternSet, err := wafregional.NewRegexPatternSet(ctx, "exampleRegexPatternSet", &wafregional.RegexPatternSetArgs{
			RegexPatternStrings: pulumi.StringArray{
				pulumi.String("one"),
				pulumi.String("two"),
			},
		})
		if err != nil {
			return err
		}
		_, err = wafregional.NewRegexMatchSet(ctx, "exampleRegexMatchSet", &wafregional.RegexMatchSetArgs{
			RegexMatchTuples: wafregional.RegexMatchSetRegexMatchTupleArray{
				&wafregional.RegexMatchSetRegexMatchTupleArgs{
					FieldToMatch: &wafregional.RegexMatchSetRegexMatchTupleFieldToMatchArgs{
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

