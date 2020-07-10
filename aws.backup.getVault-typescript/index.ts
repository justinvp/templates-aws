import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.backup.getVault({
    name: "example_backup_vault",
}, { async: true }));

