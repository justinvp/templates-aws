import pulumi
import pulumi_aws as aws

examplebucket = aws.s3.Bucket("examplebucket",
    acl="private",
    object_lock_configuration={
        "objectLockEnabled": "Enabled",
    },
    versioning={
        "enabled": True,
    })
examplebucket_object = aws.s3.BucketObject("examplebucketObject",
    bucket=examplebucket.id,
    force_destroy=True,
    key="someobject",
    object_lock_legal_hold_status="ON",
    object_lock_mode="GOVERNANCE",
    object_lock_retain_until_date="2021-12-31T23:59:60Z",
    source=pulumi.FileAsset("important.txt"))

