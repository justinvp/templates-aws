import pulumi
import pulumi_aws as aws

foo = aws.ec2.SpotFleetRequest("foo",
    iam_fleet_role="arn:aws:iam::12345678:role/spot-fleet",
    launch_specifications=[
        {
            "ami": "ami-d06a90b0",
            "availability_zone": "us-west-2a",
            "instance_type": "m1.small",
            "key_name": "my-key",
        },
        {
            "ami": "ami-d06a90b0",
            "availability_zone": "us-west-2a",
            "instance_type": "m5.large",
            "key_name": "my-key",
        },
    ],
    spot_price="0.005",
    target_capacity=2,
    valid_until="2019-11-04T20:44:20Z")

