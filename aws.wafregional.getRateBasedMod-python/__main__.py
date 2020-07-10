import pulumi
import pulumi_aws as aws

example = aws.wafregional.get_rate_based_mod(name="tfWAFRegionalRateBasedRule")

