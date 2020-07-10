import pulumi
import pulumi_aws as aws

example = aws.lambda_.ProvisionedConcurrencyConfig("example",
    function_name=aws_lambda_alias["example"]["function_name"],
    provisioned_concurrent_executions=1,
    qualifier=aws_lambda_alias["example"]["name"])

