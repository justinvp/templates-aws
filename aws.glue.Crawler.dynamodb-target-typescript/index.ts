import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Crawler("example", {
    databaseName: aws_glue_catalog_database_example.name,
    dynamodbTargets: [{
        path: "table-name",
    }],
    role: aws_iam_role_example.arn,
});

