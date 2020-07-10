import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apigatewayv2.Authorizer("example", {
    apiId: aws_apigatewayv2_api_example.id,
    authorizerType: "REQUEST",
    authorizerUri: aws_lambda_function_example.invokeArn,
    identitySources: ["route.request.header.Auth"],
});

