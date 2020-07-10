import pulumi
import pulumi_aws as aws

test_lambda_alias = aws.lambda_.Alias("testLambdaAlias",
    description="a sample description",
    function_name=aws_lambda_function["lambda_function_test"]["arn"],
    function_version="1",
    routing_config={
        "additionalVersionWeights": {
            "2": 0.5,
        },
    })

