import pulumi
import pulumi_aws as aws

example = aws.apigatewayv2.Stage("example", api_id=aws_apigatewayv2_api["example"]["id"])

