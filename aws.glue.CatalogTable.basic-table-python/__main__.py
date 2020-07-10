import pulumi
import pulumi_aws as aws

aws_glue_catalog_table = aws.glue.CatalogTable("awsGlueCatalogTable",
    database_name="MyCatalogDatabase",
    name="MyCatalogTable")

