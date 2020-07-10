import pulumi
import pulumi_aws as aws

example = aws.wafregional.get_web_acl(name="tfWAFRegionalWebACL")

