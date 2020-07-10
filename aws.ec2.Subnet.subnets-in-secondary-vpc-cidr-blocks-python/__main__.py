import pulumi
import pulumi_aws as aws

secondary_cidr = aws.ec2.VpcIpv4CidrBlockAssociation("secondaryCidr",
    cidr_block="172.2.0.0/16",
    vpc_id=aws_vpc["main"]["id"])
in_secondary_cidr = aws.ec2.Subnet("inSecondaryCidr",
    cidr_block="172.2.0.0/24",
    vpc_id=secondary_cidr.vpc_id)

