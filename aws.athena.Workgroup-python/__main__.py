import pulumi
import pulumi_aws as aws

example = aws.athena.Workgroup("example", configuration={
    "enforceWorkgroupConfiguration": True,
    "publishCloudwatchMetricsEnabled": True,
    "resultConfiguration": {
        "encryption_configuration": {
            "encryptionOption": "SSE_KMS",
            "kms_key_arn": aws_kms_key["example"]["arn"],
        },
        "output_location": "s3://{aws_s3_bucket.example.bucket}/output/",
    },
})

