import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apigatewayv2.Stage("example", {
    apiId: aws_apigatewayv2_api_example.id,
});

