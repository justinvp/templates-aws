import pulumi
import pulumi_aws as aws

# Request a Spot fleet
cheap_compute = aws.ec2.SpotFleetRequest("cheapCompute",
    allocation_strategy="diversified",
    iam_fleet_role="arn:aws:iam::12345678:role/spot-fleet",
    launch_specifications=[
        {
            "ami": "ami-1234",
            "iamInstanceProfileArn": aws_iam_instance_profile["example"]["arn"],
            "instance_type": "m4.10xlarge",
            "placement_tenancy": "dedicated",
            "spot_price": "2.793",
        },
        {
            "ami": "ami-5678",
            "availability_zone": "us-west-1a",
            "iamInstanceProfileArn": aws_iam_instance_profile["example"]["arn"],
            "instance_type": "m4.4xlarge",
            "key_name": "my-key",
            "root_block_device": [{
                "volume_size": "300",
                "volumeType": "gp2",
            }],
            "spot_price": "1.117",
            "subnet_id": "subnet-1234",
            "tags": {
                "Name": "spot-fleet-example",
            },
            "weightedCapacity": 35,
        },
    ],
    spot_price="0.03",
    target_capacity=6,
    valid_until="2019-11-04T20:44:20Z")

