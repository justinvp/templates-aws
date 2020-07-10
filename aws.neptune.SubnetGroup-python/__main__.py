import pulumi
import pulumi_aws as aws

default = aws.neptune.SubnetGroup("default",
    subnet_ids=[
        aws_subnet["frontend"]["id"],
        aws_subnet["backend"]["id"],
    ],
    tags={
        "Name": "My neptune subnet group",
    })

