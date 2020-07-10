import pulumi
import pulumi_aws as aws

example = aws.ses.EmailIdentity("example", email="email@example.com")

