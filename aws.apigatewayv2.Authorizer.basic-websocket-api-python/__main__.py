import pulumi
import pulumi_aws as aws

example = aws.apigatewayv2.Authorizer("example",
    api_id=aws_apigatewayv2_api["example"]["id"],
    authorizer_type="REQUEST",
    authorizer_uri=aws_lambda_function["example"]["invoke_arn"],
    identity_sources=["route.request.header.Auth"])

