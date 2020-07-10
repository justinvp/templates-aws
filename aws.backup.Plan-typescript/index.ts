import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.backup.Plan("example", {
    rules: [{
        ruleName: "tf_example_backup_rule",
        schedule: "cron(0 12 * * ? *)",
        targetVaultName: aws_backup_vault_test.name,
    }],
});

