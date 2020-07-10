import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const awsGlueCatalogTable = new aws.glue.CatalogTable("aws_glue_catalog_table", {
    databaseName: "MyCatalogDatabase",
    name: "MyCatalogTable",
});

