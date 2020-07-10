package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/waf"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := waf.NewSqlInjectionMatchSet(ctx, "sqlInjectionMatchSet", &waf.SqlInjectionMatchSetArgs{
			SqlInjectionMatchTuples: waf.SqlInjectionMatchSetSqlInjectionMatchTupleArray{
				&waf.SqlInjectionMatchSetSqlInjectionMatchTupleArgs{
					FieldToMatch: &waf.SqlInjectionMatchSetSqlInjectionMatchTupleFieldToMatchArgs{
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

