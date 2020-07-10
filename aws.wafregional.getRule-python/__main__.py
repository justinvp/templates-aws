import pulumi
import pulumi_aws as aws

example = aws.wafregional.get_rule(name="tfWAFRegionalRule")

