import pulumi
import pulumi_aws as aws

test = aws.ec2.NetworkInterfaceAttachment("test",
    device_index=0,
    instance_id=aws_instance["test"]["id"],
    network_interface_id=aws_network_interface["test"]["id"])

