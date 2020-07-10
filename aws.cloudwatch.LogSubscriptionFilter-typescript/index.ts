import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testLambdafunctionLogfilter = new aws.cloudwatch.LogSubscriptionFilter("test_lambdafunction_logfilter", {
    destinationArn: aws_kinesis_stream_test_logstream.arn,
    distribution: "Random",
    filterPattern: "logtype test",
    logGroup: "/aws/lambda/example_lambda_name",
    roleArn: aws_iam_role_iam_for_lambda.arn,
});

