import pulumi
import pulumi_aws as aws

example = aws.apigatewayv2.Integration("example",
    api_id=aws_apigatewayv2_api["example"]["id"],
    integration_type="MOCK")

