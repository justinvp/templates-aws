import pulumi
import pulumi_aws as aws

example = aws.datasync.Agent("example", ip_address="1.2.3.4")

