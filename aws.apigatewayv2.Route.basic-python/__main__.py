import pulumi
import pulumi_aws as aws

example = aws.apigatewayv2.Route("example",
    api_id=aws_apigatewayv2_api["example"]["id"],
    route_key="$default")

