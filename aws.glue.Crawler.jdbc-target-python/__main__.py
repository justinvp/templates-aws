import pulumi
import pulumi_aws as aws

example = aws.glue.Crawler("example",
    database_name=aws_glue_catalog_database["example"]["name"],
    jdbc_targets=[{
        "connectionName": aws_glue_connection["example"]["name"],
        "path": "database-name/%",
    }],
    role=aws_iam_role["example"]["arn"])

