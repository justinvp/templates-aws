import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const pool = new aws.cognito.UserPool("pool", {});
const resource = new aws.cognito.ResourceServer("resource", {
    identifier: "https://example.com",
    userPoolId: pool.id,
});

