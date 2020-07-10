import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2transitgateway.VpcAttachmentAccepter("example", {
    tags: {
        Name: "Example cross-account attachment",
    },
    transitGatewayAttachmentId: aws_ec2_transit_gateway_vpc_attachment_example.id,
});

