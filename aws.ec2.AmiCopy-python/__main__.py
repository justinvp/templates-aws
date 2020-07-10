import pulumi
import pulumi_aws as aws

example = aws.ec2.AmiCopy("example",
    description="A copy of ami-xxxxxxxx",
    source_ami_id="ami-xxxxxxxx",
    source_ami_region="us-west-1",
    tags={
        "Name": "HelloWorld",
    })

