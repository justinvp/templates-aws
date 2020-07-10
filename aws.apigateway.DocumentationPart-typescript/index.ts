import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleRestApi = new aws.apigateway.RestApi("example", {});
const exampleDocumentationPart = new aws.apigateway.DocumentationPart("example", {
    location: {
        method: "GET",
        path: "/example",
        type: "METHOD",
    },
    properties: "{\"description\":\"Example description\"}",
    restApiId: exampleRestApi.id,
});

