import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myapi = new aws.apigateway.RestApi("myapi", {});
const dev = new aws.apigateway.Deployment("dev", {
    restApi: myapi.id,
    stageName: "dev",
});
const prod = new aws.apigateway.Deployment("prod", {
    restApi: myapi.id,
    stageName: "prod",
});
const myUsagePlan = new aws.apigateway.UsagePlan("MyUsagePlan", {
    apiStages: [
        {
            apiId: myapi.id,
            stage: dev.stageName,
        },
        {
            apiId: myapi.id,
            stage: prod.stageName,
        },
    ],
    description: "my description",
    productCode: "MYCODE",
    quotaSettings: {
        limit: 20,
        offset: 2,
        period: "WEEK",
    },
    throttleSettings: {
        burstLimit: 5,
        rateLimit: 10,
    },
});

