package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewCatalogTable(ctx, "awsGlueCatalogTable", &glue.CatalogTableArgs{
			DatabaseName: pulumi.String("MyCatalogDatabase"),
			Name:         pulumi.String("MyCatalogTable"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

