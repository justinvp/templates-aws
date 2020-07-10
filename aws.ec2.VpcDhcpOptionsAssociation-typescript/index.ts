import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const dnsResolver = new aws.ec2.VpcDhcpOptionsAssociation("dns_resolver", {
    dhcpOptionsId: aws_vpc_dhcp_options_foo.id,
    vpcId: aws_vpc_foo.id,
});

