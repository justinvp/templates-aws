import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.apigateway.RestApi("main", {});
const test = new aws.apigateway.Response("test", {
    responseParameters: {
        "gatewayresponse.header.Authorization": "'Basic'",
    },
    responseTemplates: {
        "application/json": "{'message':$context.error.messageString}",
    },
    responseType: "UNAUTHORIZED",
    restApiId: main.id,
    statusCode: "401",
});

