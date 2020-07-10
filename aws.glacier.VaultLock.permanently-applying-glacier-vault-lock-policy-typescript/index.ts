import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glacier.VaultLock("example", {
    completeLock: true,
    policy: aws_iam_policy_document_example.json,
    vaultName: aws_glacier_vault_example.name,
});

