import pulumi
import pulumi_aws as aws

example = aws.glue.Crawler("example",
    database_name=aws_glue_catalog_database["example"]["name"],
    role=aws_iam_role["example"]["arn"],
    s3_targets=[{
        "path": f"s3://{aws_s3_bucket['example']['bucket']}",
    }])

