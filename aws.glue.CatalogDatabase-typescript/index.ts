import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const awsGlueCatalogDatabase = new aws.glue.CatalogDatabase("aws_glue_catalog_database", {
    name: "MyCatalogDatabase",
});

