import pulumi
import pulumi_aws as aws

example = aws.inspector.ResourceGroup("example", tags={
    "Env": "bar",
    "Name": "foo",
})

