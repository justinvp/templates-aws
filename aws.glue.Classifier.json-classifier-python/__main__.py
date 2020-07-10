import pulumi
import pulumi_aws as aws

example = aws.glue.Classifier("example", json_classifier={
    "jsonPath": "example",
})

