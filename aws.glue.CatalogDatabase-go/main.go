package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewCatalogDatabase(ctx, "awsGlueCatalogDatabase", &glue.CatalogDatabaseArgs{
			Name: pulumi.String("MyCatalogDatabase"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

