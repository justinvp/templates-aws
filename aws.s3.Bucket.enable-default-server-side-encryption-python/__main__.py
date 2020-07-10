import pulumi
import pulumi_aws as aws

mykey = aws.kms.Key("mykey",
    deletion_window_in_days=10,
    description="This key is used to encrypt bucket objects")
mybucket = aws.s3.Bucket("mybucket", server_side_encryption_configuration={
    "rule": {
        "applyServerSideEncryptionByDefault": {
            "kms_master_key_id": mykey.arn,
            "sseAlgorithm": "aws:kms",
        },
    },
})

