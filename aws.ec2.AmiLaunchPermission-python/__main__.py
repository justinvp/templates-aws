import pulumi
import pulumi_aws as aws

example = aws.ec2.AmiLaunchPermission("example",
    account_id="123456789012",
    image_id="ami-12345678")

