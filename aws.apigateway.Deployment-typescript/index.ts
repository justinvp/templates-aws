import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myDemoAPI = new aws.apigateway.RestApi("myDemoAPI", {description: "This is my API for demonstration purposes"});
const myDemoResource = new aws.apigateway.Resource("myDemoResource", {
    restApi: myDemoAPI.id,
    parentId: myDemoAPI.rootResourceId,
    pathPart: "test",
});
const myDemoMethod = new aws.apigateway.Method("myDemoMethod", {
    restApi: myDemoAPI.id,
    resourceId: myDemoResource.id,
    httpMethod: "GET",
    authorization: "NONE",
});
const myDemoIntegration = new aws.apigateway.Integration("myDemoIntegration", {
    restApi: myDemoAPI.id,
    resourceId: myDemoResource.id,
    httpMethod: myDemoMethod.httpMethod,
    type: "MOCK",
});
const myDemoDeployment = new aws.apigateway.Deployment("myDemoDeployment", {
    restApi: myDemoAPI.id,
    stageName: "test",
    variables: {
        answer: "42",
    },
}, {
    dependsOn: [myDemoIntegration],
});

