import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const functionName = config.require("functionName");

const existing = pulumi.output(aws.lambda.getFunction({
    functionName: functionName,
}, { async: true }));

