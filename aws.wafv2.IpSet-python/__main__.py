import pulumi
import pulumi_aws as aws

example = aws.wafv2.IpSet("example",
    addresses=[
        "1.2.3.4/32",
        "5.6.7.8/32",
    ],
    description="Example IP set",
    ip_address_version="IPV4",
    scope="REGIONAL",
    tags={
        "Tag1": "Value1",
        "Tag2": "Value2",
    })

