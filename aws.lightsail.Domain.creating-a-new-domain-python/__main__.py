import pulumi
import pulumi_aws as aws

domain_test = aws.lightsail.Domain("domainTest", domain_name="mydomain.com")

