import pulumi
import pulumi_aws as aws

current = aws.get_partition()
s3_policy = aws.iam.get_policy_document(statements=[{
    "actions": ["s3:ListBucket"],
    "resources": [f"arn:{current.partition}:s3:::my-bucket"],
    "sid": "1",
}])

