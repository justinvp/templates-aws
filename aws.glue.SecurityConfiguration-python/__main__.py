import pulumi
import pulumi_aws as aws

example = aws.glue.SecurityConfiguration("example", encryption_configuration={
    "cloudwatchEncryption": {
        "cloudwatchEncryptionMode": "DISABLED",
    },
    "jobBookmarksEncryption": {
        "jobBookmarksEncryptionMode": "DISABLED",
    },
    "s3Encryption": {
        "kms_key_arn": data["aws_kms_key"]["example"]["arn"],
        "s3EncryptionMode": "SSE-KMS",
    },
})

