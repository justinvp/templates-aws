import pulumi
import pulumi_aws as aws

xss_match_set = aws.waf.XssMatchSet("xssMatchSet", xss_match_tuples=[
    {
        "fieldToMatch": {
            "type": "URI",
        },
        "textTransformation": "NONE",
    },
    {
        "fieldToMatch": {
            "type": "QUERY_STRING",
        },
        "textTransformation": "NONE",
    },
])

