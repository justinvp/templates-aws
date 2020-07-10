import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.apigateway.RestApi("test", {});
const myusageplan = new aws.apigateway.UsagePlan("myusageplan", {
    apiStages: [{
        apiId: test.id,
        stage: aws_api_gateway_deployment_foo.stageName,
    }],
});
const mykey = new aws.apigateway.ApiKey("mykey", {});
const main = new aws.apigateway.UsagePlanKey("main", {
    keyId: mykey.id,
    keyType: "API_KEY",
    usagePlanId: myusageplan.id,
});

