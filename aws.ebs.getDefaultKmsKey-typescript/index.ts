import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = pulumi.output(aws.ebs.getDefaultKmsKey({ async: true }));
const example = new aws.ebs.Volume("example", {
    availabilityZone: "us-west-2a",
    encrypted: true,
    kmsKeyId: current.keyArn,
});

