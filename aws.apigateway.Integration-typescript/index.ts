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
    cacheKeyParameters: ["method.request.path.param"],
    cacheNamespace: "foobar",
    httpMethod: myDemoMethod.httpMethod,
    requestParameters: {
        "integration.request.header.X-Authorization": "'static'",
    },
    // Transforms the incoming XML request to JSON
    requestTemplates: {
        "application/xml": `{
   "body" : $input.json('$')
}
`,
    },
    resourceId: myDemoResource.id,
    restApi: myDemoAPI.id,
    timeoutMilliseconds: 29000,
    type: "MOCK",
});

