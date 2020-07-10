import pulumi
import pulumi_aws as aws

example = aws.ec2.get_vpc_dhcp_options(filters=[
    {
        "name": "key",
        "values": ["domain-name"],
    },
    {
        "name": "value",
        "values": ["example.com"],
    },
])

