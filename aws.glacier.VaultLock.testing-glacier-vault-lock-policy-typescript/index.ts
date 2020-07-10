import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleVault = new aws.glacier.Vault("example", {});
const examplePolicyDocument = exampleVault.arn.apply(arn => aws.iam.getPolicyDocument({
    statements: [{
        actions: ["glacier:DeleteArchive"],
        conditions: [{
            test: "NumericLessThanEquals",
            values: ["365"],
            variable: "glacier:ArchiveAgeinDays",
        }],
        effect: "Deny",
        resources: [arn],
    }],
}, { async: true }));
const exampleVaultLock = new aws.glacier.VaultLock("example", {
    completeLock: false,
    policy: examplePolicyDocument.json,
    vaultName: exampleVault.name,
});

