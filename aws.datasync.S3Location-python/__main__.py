import pulumi
import pulumi_aws as aws

example = aws.datasync.S3Location("example",
    s3_bucket_arn=aws_s3_bucket["example"]["arn"],
    s3_config={
        "bucketAccessRoleArn": aws_iam_role["example"]["arn"],
    },
    subdirectory="/example/prefix")

