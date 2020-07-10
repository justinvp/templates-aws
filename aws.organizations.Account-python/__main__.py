import pulumi
import pulumi_aws as aws

account = aws.organizations.Account("account", email="john@doe.org")

