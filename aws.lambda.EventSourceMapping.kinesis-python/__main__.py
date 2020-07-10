import pulumi
import pulumi_aws as aws

example = aws.lambda_.EventSourceMapping("example",
    event_source_arn=aws_kinesis_stream["example"]["arn"],
    function_name=aws_lambda_function["example"]["arn"],
    starting_position="LATEST")

