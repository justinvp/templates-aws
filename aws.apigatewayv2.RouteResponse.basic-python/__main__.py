import pulumi
import pulumi_aws as aws

example = aws.apigatewayv2.RouteResponse("example",
    api_id=aws_apigatewayv2_api["example"]["id"],
    route_id=aws_apigatewayv2_route["example"]["id"],
    route_response_key="$default")

