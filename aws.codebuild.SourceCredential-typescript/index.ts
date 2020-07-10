import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.codebuild.SourceCredential("example", {
    authType: "PERSONAL_ACCESS_TOKEN",
    serverType: "GITHUB",
    token: "example",
});

