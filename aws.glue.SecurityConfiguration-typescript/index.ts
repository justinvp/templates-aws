import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.SecurityConfiguration("example", {
    encryptionConfiguration: {
        cloudwatchEncryption: {
            cloudwatchEncryptionMode: "DISABLED",
        },
        jobBookmarksEncryption: {
            jobBookmarksEncryptionMode: "DISABLED",
        },
        s3Encryption: {
            kmsKeyArn: aws_kms_key_example.arn,
            s3EncryptionMode: "SSE-KMS",
        },
    },
});

