import pulumi
import pulumi_aws as aws

example = aws.glue.Classifier("example", grok_classifier={
    "classification": "example",
    "grokPattern": "example",
})

