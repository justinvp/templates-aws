import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.VpcAttachmentAccepter("example",
    tags={
        "Name": "Example cross-account attachment",
    },
    transit_gateway_attachment_id=aws_ec2_transit_gateway_vpc_attachment["example"]["id"])

