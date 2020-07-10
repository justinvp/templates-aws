import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apigatewayv2.RouteResponse("example", {
    apiId: aws_apigatewayv2_api_example.id,
    routeId: aws_apigatewayv2_route_example.id,
    routeResponseKey: "$default",
});

