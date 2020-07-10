import pulumi
import pulumi_aws as aws

example = aws.wafv2.get_regex_pattern_set(name="some-regex-pattern-set",
    scope="REGIONAL")

