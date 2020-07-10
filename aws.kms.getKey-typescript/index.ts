import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = pulumi.output(aws.kms.getKey({
    keyId: "arn:aws:kms:us-east-1:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab",
}, { async: true }));

