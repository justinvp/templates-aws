import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Crawler("example", {
    databaseName: aws_glue_catalog_database_example.name,
    jdbcTargets: [{
        connectionName: aws_glue_connection_example.name,
        path: "database-name/%",
    }],
    role: aws_iam_role_example.arn,
});

