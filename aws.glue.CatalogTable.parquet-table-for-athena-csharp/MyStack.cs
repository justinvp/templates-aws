using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var awsGlueCatalogTable = new Aws.Glue.CatalogTable("awsGlueCatalogTable", new Aws.Glue.CatalogTableArgs
        {
            DatabaseName = "MyCatalogDatabase",
            Name = "MyCatalogTable",
            Parameters = 
            {
                { "EXTERNAL", "TRUE" },
                { "parquet.compression", "SNAPPY" },
            },
            StorageDescriptor = new Aws.Glue.Inputs.CatalogTableStorageDescriptorArgs
            {
                Columns = 
                {
                    new Aws.Glue.Inputs.CatalogTableStorageDescriptorColumnArgs
                    {
                        Name = "my_string",
                        Type = "string",
                    },
                    new Aws.Glue.Inputs.CatalogTableStorageDescriptorColumnArgs
                    {
                        Name = "my_double",
                        Type = "double",
                    },
                    new Aws.Glue.Inputs.CatalogTableStorageDescriptorColumnArgs
                    {
                        Comment = "",
                        Name = "my_date",
                        Type = "date",
                    },
                    new Aws.Glue.Inputs.CatalogTableStorageDescriptorColumnArgs
                    {
                        Comment = "",
                        Name = "my_bigint",
                        Type = "bigint",
                    },
                    new Aws.Glue.Inputs.CatalogTableStorageDescriptorColumnArgs
                    {
                        Comment = "",
                        Name = "my_struct",
                        Type = "struct<my_nested_string:string>",
                    },
                },
                InputFormat = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat",
                Location = "s3://my-bucket/event-streams/my-stream",
                OutputFormat = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat",
                SerDeInfo = new Aws.Glue.Inputs.CatalogTableStorageDescriptorSerDeInfoArgs
                {
                    Name = "my-stream",
                    Parameters = 
                    {
                        { "serialization.format", "1" },
                    },
                    SerializationLibrary = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe",
                },
            },
            TableType = "EXTERNAL_TABLE",
        });
    }

}

