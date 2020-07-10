import pulumi
import pulumi_aws as aws

example = aws.ram.ResourceShare("example",
    allow_external_principals=True,
    tags={
        "Environment": "Production",
    })

