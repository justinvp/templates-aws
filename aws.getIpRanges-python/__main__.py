import pulumi
import pulumi_aws as aws

european_ec2 = aws.get_ip_ranges(regions=[
        "eu-west-1",
        "eu-central-1",
    ],
    services=["ec2"])
from_europe = aws.ec2.SecurityGroup("fromEurope",
    ingress=[{
        "from_port": "443",
        "to_port": "443",
        "protocol": "tcp",
        "cidr_blocks": european_ec2.cidr_blocks,
        "ipv6_cidr_blocks": european_ec2.ipv6_cidr_blocks,
    }],
    tags={
        "CreateDate": european_ec2.create_date,
        "SyncToken": european_ec2.sync_token,
    })

