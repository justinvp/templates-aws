import pulumi
import pulumi_aws as aws

example = aws.glue.Classifier("example", csv_classifier={
    "allowSingleColumn": False,
    "containsHeader": "PRESENT",
    "delimiter": ",",
    "disableValueTrimming": False,
    "header": [
        "example1",
        "example2",
    ],
    "quoteSymbol": "'",
})

