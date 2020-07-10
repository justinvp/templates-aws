import pulumi
import pulumi_aws as aws

example = aws.glue.Classifier("example", xml_classifier={
    "classification": "example",
    "rowTag": "example",
})

