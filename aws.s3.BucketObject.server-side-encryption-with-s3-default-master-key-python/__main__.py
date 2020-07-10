import pulumi
import pulumi_aws as aws

examplebucket = aws.s3.Bucket("examplebucket", acl="private")
examplebucket_object = aws.s3.BucketObject("examplebucketObject",
    bucket=examplebucket.id,
    key="someobject",
    server_side_encryption="aws:kms",
    source=pulumi.FileAsset("index.html"))

