import pulumi
import pulumi_aws as aws

test_destination = aws.cloudwatch.LogDestination("testDestination",
    role_arn=aws_iam_role["iam_for_cloudwatch"]["arn"],
    target_arn=aws_kinesis_stream["kinesis_for_cloudwatch"]["arn"])

