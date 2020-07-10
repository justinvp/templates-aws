import pulumi
import pulumi_aws as aws

sn_ec2 = aws.ec2.VpcEndpointSubnetAssociation("snEc2",
    subnet_id=aws_subnet["sn"]["id"],
    vpc_endpoint_id=aws_vpc_endpoint["ec2"]["id"])

