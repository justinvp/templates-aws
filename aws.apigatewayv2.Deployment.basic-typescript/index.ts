import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apigatewayv2.Deployment("example", {
    apiId: aws_apigatewayv2_route_example.apiId,
    description: "Example deployment",
});

