import pulumi
import pulumi_aws as aws

aws_glue_catalog_database = aws.glue.CatalogDatabase("awsGlueCatalogDatabase", name="MyCatalogDatabase")

