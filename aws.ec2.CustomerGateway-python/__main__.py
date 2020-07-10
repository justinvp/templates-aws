import pulumi
import pulumi_aws as aws

main = aws.ec2.CustomerGateway("main",
    bgp_asn=65000,
    ip_address="172.83.124.10",
    tags={
        "Name": "main-customer-gateway",
    },
    type="ipsec.1")

