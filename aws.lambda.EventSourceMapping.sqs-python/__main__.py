import pulumi
import pulumi_aws as aws

example = aws.lambda_.EventSourceMapping("example",
    event_source_arn=aws_sqs_queue["sqs_queue_test"]["arn"],
    function_name=aws_lambda_function["example"]["arn"])

