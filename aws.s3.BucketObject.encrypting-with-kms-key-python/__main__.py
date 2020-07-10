import pulumi
import pulumi_aws as aws

examplekms = aws.kms.Key("examplekms",
    deletion_window_in_days=7,
    description="KMS key 1")
examplebucket = aws.s3.Bucket("examplebucket", acl="private")
examplebucket_object = aws.s3.BucketObject("examplebucketObject",
    bucket=examplebucket.id,
    key="someobject",
    kms_key_id=examplekms.arn,
    source=pulumi.FileAsset("index.html"))

