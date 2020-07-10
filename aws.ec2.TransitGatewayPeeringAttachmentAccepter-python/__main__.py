import pulumi
import pulumi_aws as aws

example = aws.ec2.TransitGatewayPeeringAttachmentAccepter("example",
    tags={
        "Name": "Example cross-account attachment",
    },
    transit_gateway_attachment_id=aws_ec2_transit_gateway_peering_attachment["example"]["id"])

