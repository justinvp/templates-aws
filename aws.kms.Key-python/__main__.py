import pulumi
import pulumi_aws as aws

key = aws.kms.Key("key",
    deletion_window_in_days=10,
    description="KMS key 1")

