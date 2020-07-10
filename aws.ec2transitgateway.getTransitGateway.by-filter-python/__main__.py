import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.get_transit_gateway(filters=[{
    "name": "options.amazon-side-asn",
    "values": ["64512"],
}])

