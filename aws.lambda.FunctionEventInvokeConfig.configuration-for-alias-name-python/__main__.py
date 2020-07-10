import pulumi
import pulumi_aws as aws

example = aws.lambda_.FunctionEventInvokeConfig("example",
    function_name=aws_lambda_alias["example"]["function_name"],
    qualifier=aws_lambda_alias["example"]["name"])
# ... other configuration ...

