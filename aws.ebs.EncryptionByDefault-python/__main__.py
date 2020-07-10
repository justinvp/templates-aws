import pulumi
import pulumi_aws as aws

example = aws.ebs.EncryptionByDefault("example", enabled=True)

