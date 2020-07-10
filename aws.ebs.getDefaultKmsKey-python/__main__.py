import pulumi
import pulumi_aws as aws

current = aws.ebs.get_default_kms_key()
example = aws.ebs.Volume("example",
    availability_zone="us-west-2a",
    encrypted=True,
    kms_key_id=current.key_arn)

