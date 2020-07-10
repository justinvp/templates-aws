import pulumi
import pulumi_aws as aws

dns_resolver = aws.ec2.VpcDhcpOptionsAssociation("dnsResolver",
    dhcp_options_id=aws_vpc_dhcp_options["foo"]["id"],
    vpc_id=aws_vpc["foo"]["id"])

