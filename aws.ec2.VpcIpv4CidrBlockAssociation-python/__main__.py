import pulumi
import pulumi_aws as aws

main = aws.ec2.Vpc("main", cidr_block="10.0.0.0/16")
secondary_cidr = aws.ec2.VpcIpv4CidrBlockAssociation("secondaryCidr",
    cidr_block="172.2.0.0/16",
    vpc_id=main.id)

