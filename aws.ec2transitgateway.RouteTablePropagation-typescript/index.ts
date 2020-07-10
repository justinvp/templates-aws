import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2transitgateway.RouteTablePropagation("example", {
    transitGatewayAttachmentId: aws_ec2_transit_gateway_vpc_attachment_example.id,
    transitGatewayRouteTableId: aws_ec2_transit_gateway_route_table_example.id,
});

