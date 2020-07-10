import pulumi
import pulumi_aws as aws

example = aws.glue.Job("example",
    command={
        "scriptLocation": f"s3://{aws_s3_bucket['example']['bucket']}/example.scala",
    },
    default_arguments={
        "--job-language": "scala",
    },
    role_arn=aws_iam_role["example"]["arn"])

