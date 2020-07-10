import pulumi
import pulumi_aws as aws

sql_injection_match_set = aws.wafregional.SqlInjectionMatchSet("sqlInjectionMatchSet", sql_injection_match_tuples=[{
    "fieldToMatch": {
        "type": "QUERY_STRING",
    },
    "textTransformation": "URL_DECODE",
}])

