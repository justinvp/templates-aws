package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/waf"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := waf.NewXssMatchSet(ctx, "xssMatchSet", &waf.XssMatchSetArgs{
			XssMatchTuples: waf.XssMatchSetXssMatchTupleArray{
				&waf.XssMatchSetXssMatchTupleArgs{
					FieldToMatch: &waf.XssMatchSetXssMatchTupleFieldToMatchArgs{
						Type: pulumi.String("URI"),
					},
					TextTransformation: pulumi.String("NONE"),
				},
				&waf.XssMatchSetXssMatchTupleArgs{
					FieldToMatch: &waf.XssMatchSetXssMatchTupleFieldToMatchArgs{
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

