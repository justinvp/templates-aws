import pulumi
import pulumi_aws as aws

example = aws.wafv2.get_web_acl(name="some-web-acl",
    scope="REGIONAL")

