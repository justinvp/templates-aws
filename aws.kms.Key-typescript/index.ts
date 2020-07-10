import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const key = new aws.kms.Key("a", {
    deletionWindowInDays: 10,
    description: "KMS key 1",
});

