import pulumi
import pulumi_aws as aws

ipset = aws.waf.IpSet("ipset", ip_set_descriptors=[
    {
        "type": "IPV4",
        "value": "192.0.7.0/24",
    },
    {
        "type": "IPV4",
        "value": "10.16.16.0/16",
    },
])

