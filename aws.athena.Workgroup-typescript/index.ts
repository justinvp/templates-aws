import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.athena.Workgroup("example", {
    configuration: {
        enforceWorkgroupConfiguration: true,
        publishCloudwatchMetricsEnabled: true,
        resultConfiguration: {
            encryptionConfiguration: {
                encryptionOption: "SSE_KMS",
                kmsKeyArn: aws_kms_key_example.arn,
            },
            outputLocation: "s3://{aws_s3_bucket.example.bucket}/output/",
        },
    },
});

