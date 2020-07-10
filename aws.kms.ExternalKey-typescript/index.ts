import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kms.ExternalKey("example", {
    description: "KMS EXTERNAL for AMI encryption",
});

