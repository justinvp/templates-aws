import pulumi
import pulumi_aws as aws

production = aws.lambda.get_alias(function_name="my-lambda-func",
    name="production")

