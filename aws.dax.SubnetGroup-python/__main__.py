import pulumi
import pulumi_aws as aws

example = aws.dax.SubnetGroup("example", subnet_ids=[
    aws_subnet["example1"]["id"],
    aws_subnet["example2"]["id"],
])

