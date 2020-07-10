package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewClassifier(ctx, "example", &glue.ClassifierArgs{
			CsvClassifier: &glue.ClassifierCsvClassifierArgs{
				AllowSingleColumn:    pulumi.Bool(false),
				ContainsHeader:       pulumi.String("PRESENT"),
				Delimiter:            pulumi.String(","),
				DisableValueTrimming: pulumi.Bool(false),
				Header: pulumi.StringArray{
					pulumi.String("example1"),
					pulumi.String("example2"),
				},
				QuoteSymbol: pulumi.String("'"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}

