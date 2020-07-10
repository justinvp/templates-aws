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
			Parameters: pulumi.StringMap{
				"EXTERNAL":            pulumi.String("TRUE"),
				"parquet.compression": pulumi.String("SNAPPY"),
			},
			StorageDescriptor: &glue.CatalogTableStorageDescriptorArgs{
				Columns: glue.CatalogTableStorageDescriptorColumnArray{
					&glue.CatalogTableStorageDescriptorColumnArgs{
						Name: pulumi.String("my_string"),
						Type: pulumi.String("string"),
					},
					&glue.CatalogTableStorageDescriptorColumnArgs{
						Name: pulumi.String("my_double"),
						Type: pulumi.String("double"),
					},
					&glue.CatalogTableStorageDescriptorColumnArgs{
						Comment: pulumi.String(""),
						Name:    pulumi.String("my_date"),
						Type:    pulumi.String("date"),
					},
					&glue.CatalogTableStorageDescriptorColumnArgs{
						Comment: pulumi.String(""),
						Name:    pulumi.String("my_bigint"),
						Type:    pulumi.String("bigint"),
					},
					&glue.CatalogTableStorageDescriptorColumnArgs{
						Comment: pulumi.String(""),
						Name:    pulumi.String("my_struct"),
						Type:    pulumi.String("struct<my_nested_string:string>"),
					},
				},
				InputFormat:  pulumi.String("org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"),
				Location:     pulumi.String("s3://my-bucket/event-streams/my-stream"),
				OutputFormat: pulumi.String("org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"),
				SerDeInfo: &glue.CatalogTableStorageDescriptorSerDeInfoArgs{
					Name: pulumi.String("my-stream"),
					Parameters: pulumi.StringMap{
						"serialization.format": pulumi.String("1"),
					},
					SerializationLibrary: pulumi.String("org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"),
				},
			},
			TableType: pulumi.String("EXTERNAL_TABLE"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}

