import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.VpnGatewayRoutePropagation("example", {
    routeTableId: aws_route_table_example.id,
    vpnGatewayId: aws_vpn_gateway_example.id,
});

