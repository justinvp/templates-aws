import pulumi
import pulumi_aws as aws

dns_resolver = aws.ec2.VpcDhcpOptions("dnsResolver", domain_name_servers=[
    "8.8.8.8",
    "8.8.4.4",
])

