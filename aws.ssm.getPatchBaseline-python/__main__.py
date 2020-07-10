import pulumi
import pulumi_aws as aws

centos = aws.ssm.get_patch_baseline(name_prefix="AWS-",
    operating_system="CENTOS",
    owner="AWS")

