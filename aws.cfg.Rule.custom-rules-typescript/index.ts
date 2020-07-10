import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleRecorder = new aws.cfg.Recorder("example", {});
const exampleFunction = new aws.lambda.Function("example", {});
const examplePermission = new aws.lambda.Permission("example", {
    action: "lambda:InvokeFunction",
    function: exampleFunction.arn,
    principal: "config.amazonaws.com",
});
const exampleRule = new aws.cfg.Rule("example", {
    source: {
        owner: "CUSTOM_LAMBDA",
        sourceIdentifier: exampleFunction.arn,
    },
}, { dependsOn: [exampleRecorder, examplePermission] });

