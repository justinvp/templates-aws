import pulumi
import pulumi_aws as aws

config = pulumi.Config()
function_name = config.require_object("functionName")
existing = aws.lambda.get_function(function_name=function_name)

