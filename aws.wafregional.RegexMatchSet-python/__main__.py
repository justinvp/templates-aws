import pulumi
import pulumi_aws as aws

example_regex_pattern_set = aws.wafregional.RegexPatternSet("exampleRegexPatternSet", regex_pattern_strings=[
    "one",
    "two",
])
example_regex_match_set = aws.wafregional.RegexMatchSet("exampleRegexMatchSet", regex_match_tuples=[{
    "fieldToMatch": {
        "data": "User-Agent",
        "type": "HEADER",
    },
    "regexPatternSetId": example_regex_pattern_set.id,
    "textTransformation": "NONE",
}])

