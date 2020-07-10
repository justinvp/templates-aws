import pulumi
import pulumi_aws as aws

ubuntu = aws.ec2.get_launch_configuration(name="test-launch-config")

