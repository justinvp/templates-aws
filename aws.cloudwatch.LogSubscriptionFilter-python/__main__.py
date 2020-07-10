import pulumi
import pulumi_aws as aws

test_lambdafunction_logfilter = aws.cloudwatch.LogSubscriptionFilter("testLambdafunctionLogfilter",
    destination_arn=aws_kinesis_stream["test_logstream"]["arn"],
    distribution="Random",
    filter_pattern="logtype test",
    log_group="/aws/lambda/example_lambda_name",
    role_arn=aws_iam_role["iam_for_lambda"]["arn"])

