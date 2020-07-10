import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apigatewayv2.Authorizer("example", {
    apiId: aws_apigatewayv2_api_example.id,
    authorizerType: "JWT",
    identitySources: ["$request.header.Authorization"],
    jwtConfiguration: {
        audiences: ["example"],
        issuer: pulumi.interpolate`https://${aws_cognito_user_pool_example.endpoint}`,
    },
});

