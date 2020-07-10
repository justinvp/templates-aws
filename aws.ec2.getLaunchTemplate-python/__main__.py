import pulumi
import pulumi_aws as aws

default = aws.ec2.get_launch_template(name="my-launch-template")

