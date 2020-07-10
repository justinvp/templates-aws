import pulumi
import pulumi_aws as aws

example = aws.ec2.get_instance_type_offerings(filters=[
        {
            "name": "instance-type",
            "values": [
                "t2.micro",
                "t3.micro",
            ],
        },
        {
            "name": "location",
            "values": ["usw2-az4"],
        },
    ],
    location_type="availability-zone-id")

