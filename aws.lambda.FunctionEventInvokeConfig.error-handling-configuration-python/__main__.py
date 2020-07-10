import pulumi
import pulumi_aws as aws

example = aws.lambda_.FunctionEventInvokeConfig("example",
    function_name=aws_lambda_alias["example"]["function_name"],
    maximum_event_age_in_seconds=60,
    maximum_retry_attempts=0)

