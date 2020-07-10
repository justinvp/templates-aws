import pulumi
import pulumi_aws as aws

example = aws.wafv2.RegexPatternSet("example",
    description="Example regex pattern set",
    regular_expressions=[
        {
            "regexString": "one",
        },
        {
            "regexString": "two",
        },
    ],
    scope="REGIONAL",
    tags={
        "Tag1": "Value1",
        "Tag2": "Value2",
    })

