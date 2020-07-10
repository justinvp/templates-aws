import pulumi
import pulumi_aws as aws

default_az1 = aws.ec2.DefaultSubnet("defaultAz1",
    availability_zone="us-west-2a",
    tags={
        "Name": "Default subnet for us-west-2a",
    })

