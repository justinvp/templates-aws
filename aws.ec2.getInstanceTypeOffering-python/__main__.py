import pulumi
import pulumi_aws as aws

example = aws.ec2.get_instance_type_offering(filters=[{
        "name": "instance-type",
        "values": [
            "t1.micro",
            "t2.micro",
            "t3.micro",
        ],
    }],
    preferred_instance_types=[
        "t3.micro",
        "t2.micro",
        "t1.micro",
    ])

