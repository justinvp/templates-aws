import pulumi
import pulumi_aws as aws

test = aws.lightsail.StaticIp("test")

