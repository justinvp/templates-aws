import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cognito.UserPool("example", {});
const main = new aws.cognito.UserPoolDomain("main", {
    domain: "example-domain",
    userPoolId: example.id,
});

