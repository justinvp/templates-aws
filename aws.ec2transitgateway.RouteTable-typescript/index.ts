import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2transitgateway.RouteTable("example", {
    transitGatewayId: aws_ec2_transit_gateway_example.id,
});

