import pulumi
import pulumi_aws as aws

example = aws.ec2.AvailabilityZoneGroup("example",
    group_name="us-west-2-lax-1",
    opt_in_status="opted-in")

