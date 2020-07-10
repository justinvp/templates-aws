import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.codebuild.SourceCredential("example", {
    authType: "BASIC_AUTH",
    serverType: "BITBUCKET",
    token: "example",
    userName: "test-user",
});

