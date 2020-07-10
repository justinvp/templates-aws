import pulumi
import pulumi_aws as aws

example = aws.waf.get_rate_based_rule(name="tfWAFRateBasedRule")

