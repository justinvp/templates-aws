import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleLayerVersion = new aws.lambda.LayerVersion("example", {});
const exampleFunction = new aws.lambda.Function("example", {
    // ... other configuration ...
    layers: [exampleLayerVersion.arn],
});

