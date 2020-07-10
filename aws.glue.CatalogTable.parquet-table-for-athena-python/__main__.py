import pulumi
import pulumi_aws as aws

aws_glue_catalog_table = aws.glue.CatalogTable("awsGlueCatalogTable",
    database_name="MyCatalogDatabase",
    name="MyCatalogTable",
    parameters={
        "EXTERNAL": "TRUE",
        "parquet.compression": "SNAPPY",
    },
    storage_descriptor={
        "columns": [
            {
                "name": "my_string",
                "type": "string",
            },
            {
                "name": "my_double",
                "type": "double",
            },
            {
                "comment": "",
                "name": "my_date",
                "type": "date",
            },
            {
                "comment": "",
                "name": "my_bigint",
                "type": "bigint",
            },
            {
                "comment": "",
                "name": "my_struct",
                "type": "struct<my_nested_string:string>",
            },
        ],
        "inputFormat": "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat",
        "location": "s3://my-bucket/event-streams/my-stream",
        "outputFormat": "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat",
        "serDeInfo": {
            "name": "my-stream",
            "parameters": {
                "serialization.format": 1,
            },
            "serializationLibrary": "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe",
        },
    },
    table_type="EXTERNAL_TABLE")

