import pulumi
import pulumi_aws as aws

example = aws.iot.Thing("example", attributes={
    "First": "examplevalue",
})

