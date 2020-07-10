import pulumi
import pulumi_aws as aws

byte_set = aws.waf.ByteMatchSet("byteSet", byte_match_tuples=[{
    "fieldToMatch": {
        "data": "referer",
        "type": "HEADER",
    },
    "positionalConstraint": "CONTAINS",
    "targetString": "badrefer1",
    "textTransformation": "NONE",
}])

