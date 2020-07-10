import pulumi
import pulumi_aws as aws

example = aws.ec2.get_local_gateway_virtual_interface_group(local_gateway_id=data["aws_ec2_local_gateway"]["example"]["id"])

