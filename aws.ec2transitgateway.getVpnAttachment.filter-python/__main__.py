import pulumi
import pulumi_aws as aws

test = aws.ec2transitgateway.get_vpn_attachment(filters=[{
    "name": "resource-id",
    "values": ["some-resource"],
}])

