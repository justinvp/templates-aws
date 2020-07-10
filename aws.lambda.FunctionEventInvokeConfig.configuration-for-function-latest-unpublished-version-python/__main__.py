import pulumi
import pulumi_aws as aws

example = aws.lambda_.FunctionEventInvokeConfig("example",
    function_name=aws_lambda_function["example"]["function_name"],
    qualifier="$LATEST")
# ... other configuration ...

