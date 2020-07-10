import pulumi
import pulumi_aws as aws

bar = aws.ec2.get_network_interface(id="eni-01234567")

