import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.get_peering_attachment(filters=[{
    "name": "transit-gateway-attachment-id",
    "values": ["tgw-attach-12345678"],
}])

