import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2transitgateway.Route("example", {
    blackhole: true,
    destinationCidrBlock: "0.0.0.0/0",
    transitGatewayRouteTableId: aws_ec2_transit_gateway_example.associationDefaultRouteTableId,
});

