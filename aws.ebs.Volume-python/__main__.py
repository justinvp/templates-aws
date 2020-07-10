import pulumi
import pulumi_aws as aws

example = aws.ebs.Volume("example",
    availability_zone="us-west-2a",
    size=40,
    tags={
        "Name": "HelloWorld",
    })

