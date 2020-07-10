import pulumi
import pulumi_aws as aws

test = aws.efs.AccessPoint("test", file_system_id=aws_efs_file_system["foo"]["id"])

