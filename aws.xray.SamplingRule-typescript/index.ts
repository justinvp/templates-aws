import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.xray.SamplingRule("example", {
    attributes: {
        Hello: "Tris",
    },
    fixedRate: 0.05,
    host: "*",
    httpMethod: "*",
    priority: 10000,
    reservoirSize: 1,
    resourceArn: "*",
    ruleName: "example",
    serviceName: "*",
    serviceType: "*",
    urlPath: "*",
    version: 1,
});

