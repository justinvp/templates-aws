import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apigatewayv2.IntegrationResponse("example", {
    apiId: aws_apigatewayv2_api_example.id,
    integrationId: aws_apigatewayv2_integration_example.id,
    integrationResponseKey: "/200/",
});

