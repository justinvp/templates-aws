import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.backup.Vault("example", {
    kmsKeyArn: aws_kms_key_example.arn,
});

