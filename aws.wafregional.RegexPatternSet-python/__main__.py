import pulumi
import pulumi_aws as aws

example = aws.wafregional.RegexPatternSet("example", regex_pattern_strings=[
    "one",
    "two",
])

