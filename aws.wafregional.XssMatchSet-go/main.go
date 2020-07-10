package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/wafregional"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := wafregional.NewXssMatchSet(ctx, "xssMatchSet", &wafregional.XssMatchSetArgs{
			XssMatchTuples: wafregional.XssMatchSetXssMatchTupleArray{
				&wafregional.XssMatchSetXssMatchTupleArgs{
					FieldToMatch: &wafregional.XssMatchSetXssMatchTupleFieldToMatchArgs{
						Type: pulumi.String("URI"),
					},
					TextTransformation: pulumi.String("NONE"),
				},
				&wafregional.XssMatchSetXssMatchTupleArgs{
					FieldToMatch: &wafregional.XssMatchSetXssMatchTupleFieldToMatchArgs{
						Type: pulumi.String("QUERY_STRING"),
					},
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

