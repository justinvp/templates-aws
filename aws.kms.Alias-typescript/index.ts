import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const key = new aws.kms.Key("a", {});
const alias = new aws.kms.Alias("a", {
    targetKeyId: key.keyId,
});

