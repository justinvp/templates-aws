import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const awsGlueCatalogTable = new aws.glue.CatalogTable("aws_glue_catalog_table", {
    databaseName: "MyCatalogDatabase",
    name: "MyCatalogTable",
    parameters: {
        EXTERNAL: "TRUE",
        "parquet.compression": "SNAPPY",
    },
    storageDescriptor: {
        columns: [
            {
                name: "my_string",
                type: "string",
            },
            {
                name: "my_double",
                type: "double",
            },
            {
                comment: "",
                name: "my_date",
                type: "date",
            },
            {
                comment: "",
                name: "my_bigint",
                type: "bigint",
            },
            {
                comment: "",
                name: "my_struct",
                type: "struct<my_nested_string:string>",
            },
        ],
        inputFormat: "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat",
        location: "s3://my-bucket/event-streams/my-stream",
        outputFormat: "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat",
        serDeInfo: {
            name: "my-stream",
            parameters: {
                "serialization.format": 1,
            },
            serializationLibrary: "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe",
        },
    },
    tableType: "EXTERNAL_TABLE",
});

