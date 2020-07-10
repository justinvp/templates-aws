import pulumi
import pulumi_aws as aws

test = aws.ec2.NetworkInterface("test",
    attachments=[{
        "device_index": 1,
        "instance": aws_instance["test"]["id"],
    }],
    private_ips=["10.0.0.50"],
    security_groups=[aws_security_group["web"]["id"]],
    subnet_id=aws_subnet["public_a"]["id"])

