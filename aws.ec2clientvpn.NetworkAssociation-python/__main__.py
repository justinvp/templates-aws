import pulumi
import pulumi_aws as aws

example = aws.ec2clientvpn.NetworkAssociation("example",
    client_vpn_endpoint_id=aws_ec2_client_vpn_endpoint["example"]["id"],
    subnet_id=aws_subnet["example"]["id"])

