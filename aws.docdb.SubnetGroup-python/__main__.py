import pulumi
import pulumi_aws as aws

default = aws.docdb.SubnetGroup("default",
    subnet_ids=[
        aws_subnet["frontend"]["id"],
        aws_subnet["backend"]["id"],
    ],
    tags={
        "Name": "My docdb subnet group",
    })

