import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apigateway.RequestValidator("example", {
    restApi: aws_api_gateway_rest_api_example.id,
    validateRequestBody: true,
    validateRequestParameters: true,
});

