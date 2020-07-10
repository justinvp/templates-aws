import pulumi
import pulumi_aws as aws

example = aws.waf.get_rule(name="tfWAFRule")

