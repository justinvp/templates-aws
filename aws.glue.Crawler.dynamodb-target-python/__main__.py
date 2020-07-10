import pulumi
import pulumi_aws as aws

example = aws.glue.Crawler("example",
    database_name=aws_glue_catalog_database["example"]["name"],
    dynamodb_targets=[{
        "path": "table-name",
    }],
    role=aws_iam_role["example"]["arn"])

