import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.s3.AccountPublicAccessBlock("example", {
    blockPublicAcls: true,
    blockPublicPolicy: true,
});

