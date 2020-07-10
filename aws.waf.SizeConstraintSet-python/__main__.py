import pulumi
import pulumi_aws as aws

size_constraint_set = aws.waf.SizeConstraintSet("sizeConstraintSet", size_constraints=[{
    "comparison_operator": "EQ",
    "fieldToMatch": {
        "type": "BODY",
    },
    "size": "4096",
    "textTransformation": "NONE",
}])

