import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testRestApi = new aws.apigateway.RestApi("test", {
    description: "This is my API for demonstration purposes",
});
const testResource = new aws.apigateway.Resource("test", {
    parentId: testRestApi.rootResourceId,
    pathPart: "mytestresource",
    restApi: testRestApi.id,
});
const testMethod = new aws.apigateway.Method("test", {
    authorization: "NONE",
    httpMethod: "GET",
    resourceId: testResource.id,
    restApi: testRestApi.id,
});
const testIntegration = new aws.apigateway.Integration("test", {
    httpMethod: testMethod.httpMethod,
    resourceId: testResource.id,
    restApi: testRestApi.id,
    type: "MOCK",
});
const testDeployment = new aws.apigateway.Deployment("test", {
    restApi: testRestApi.id,
    stageName: "dev",
}, { dependsOn: [testIntegration] });
const testStage = new aws.apigateway.Stage("test", {
    deployment: testDeployment.id,
    restApi: testRestApi.id,
    stageName: "prod",
});
const methodSettings = new aws.apigateway.MethodSettings("s", {
    methodPath: pulumi.interpolate`${testResource.pathPart}/${testMethod.httpMethod}`,
    restApi: testRestApi.id,
    settings: {
        loggingLevel: "INFO",
        metricsEnabled: true,
    },
    stageName: testStage.stageName,
});

