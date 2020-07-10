import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleRestApi = new aws.apigateway.RestApi("example", {});
const exampleDocumentationPart = new aws.apigateway.DocumentationPart("example", {
    location: {
        type: "API",
    },
    properties: "{\"description\":\"Example\"}",
    restApiId: exampleRestApi.id,
});
const exampleDocumentationVersion = new aws.apigateway.DocumentationVersion("example", {
    description: "Example description",
    restApiId: exampleRestApi.id,
    version: "example_version",
}, { dependsOn: [exampleDocumentationPart] });

