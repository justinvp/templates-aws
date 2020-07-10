package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/wafregional"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := wafregional.NewSqlInjectionMatchSet(ctx, "sqlInjectionMatchSet", &wafregional.SqlInjectionMatchSetArgs{
			SqlInjectionMatchTuples: wafregional.SqlInjectionMatchSetSqlInjectionMatchTupleArray{
				&wafregional.SqlInjectionMatchSetSqlInjectionMatchTupleArgs{
					FieldToMatch: &wafregional.SqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatchArgs{
						Type: pulumi.String("QUERY_STRING"),
					},
					TextTransformation: pulumi.String("URL_DECODE"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

