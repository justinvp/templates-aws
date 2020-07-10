import pulumi
import pulumi_aws as aws

example = aws.waf.get_web_acl(name="tfWAFWebACL")

