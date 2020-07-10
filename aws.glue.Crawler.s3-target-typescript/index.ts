import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Crawler("example", {
    databaseName: aws_glue_catalog_database_example.name,
    role: aws_iam_role_example.arn,
    s3Targets: [{
        path: pulumi.interpolate`s3://${aws_s3_bucket_example.bucket}`,
    }],
});

