import pulumi
import pulumi_aws as aws

example = aws.waf.RegexPatternSet("example", regex_pattern_strings=[
    "one",
    "two",
])

