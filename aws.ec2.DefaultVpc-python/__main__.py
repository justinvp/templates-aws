import pulumi
import pulumi_aws as aws

default = aws.ec2.DefaultVpc("default", tags={
    "Name": "Default VPC",
})

