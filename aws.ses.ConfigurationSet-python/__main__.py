import pulumi
import pulumi_aws as aws

test = aws.ses.ConfigurationSet("test")

