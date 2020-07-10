import pulumi
import pulumi_aws as aws

example = aws.apigatewayv2.Api("example",
    protocol_type="WEBSOCKET",
    route_selection_expression="$request.body.action")

