import pulumi
import pulumi_aws as aws

default = aws.ec2.CapacityReservation("default",
    availability_zone="eu-west-1a",
    instance_count=1,
    instance_platform="Linux/UNIX",
    instance_type="t2.micro")

