import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testDestination = new aws.cloudwatch.LogDestination("test_destination", {
    roleArn: aws_iam_role_iam_for_cloudwatch.arn,
    targetArn: aws_kinesis_stream_kinesis_for_cloudwatch.arn,
});

