import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myDemoAPI = new aws.apigateway.RestApi("MyDemoAPI", {
    description: "This is my API for demonstration purposes",
});
const myDemoResource = new aws.apigateway.Resource("MyDemoResource", {
    parentId: myDemoAPI.rootResourceId,
    pathPart: "mydemoresource",
    restApi: myDemoAPI.id,
});
const myDemoMethod = new aws.apigateway.Method("MyDemoMethod", {
    authorization: "NONE",
    httpMethod: "GET",
    resourceId: myDemoResource.id,
    restApi: myDemoAPI.id,
});
const myDemoIntegration = new aws.apigateway.Integration("MyDemoIntegration", {
    httpMethod: myDemoMethod.httpMethod,
    resourceId: myDemoResource.id,
    restApi: myDemoAPI.id,
    type: "MOCK",
});
const response200 = new aws.apigateway.MethodResponse("response_200", {
    httpMethod: myDemoMethod.httpMethod,
    resourceId: myDemoResource.id,
    restApi: myDemoAPI.id,
    statusCode: "200",
});
const myDemoIntegrationResponse = new aws.apigateway.IntegrationResponse("MyDemoIntegrationResponse", {
    httpMethod: myDemoMethod.httpMethod,
    resourceId: myDemoResource.id,
    // Transforms the backend JSON response to XML
    responseTemplates: {
        "application/xml": `#set($inputRoot = $input.path('$'))
<?xml version="1.0" encoding="UTF-8"?>
<message>
    $inputRoot.body
</message>
`,
    },
    restApi: myDemoAPI.id,
    statusCode: response200.statusCode,
});

