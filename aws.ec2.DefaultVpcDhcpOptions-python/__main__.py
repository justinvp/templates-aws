import pulumi
import pulumi_aws as aws

default = aws.ec2.DefaultVpcDhcpOptions("default", tags={
    "Name": "Default DHCP Option Set",
})

