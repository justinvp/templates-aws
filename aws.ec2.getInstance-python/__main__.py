import pulumi
import pulumi_aws as aws

foo = aws.ec2.get_instance(filters=[
        {
            "name": "image-id",
            "values": ["ami-xxxxxxxx"],
        },
        {
            "name": "tag:Name",
            "values": ["instance-name-tag"],
        },
    ],
    instance_id="i-instanceid")

