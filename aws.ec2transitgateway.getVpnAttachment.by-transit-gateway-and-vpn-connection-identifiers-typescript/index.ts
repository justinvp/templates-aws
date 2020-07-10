import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.all([aws_ec2_transit_gateway_example.id, aws_vpn_connection_example.id]).apply(([aws_ec2_transit_gateway_exampleId, aws_vpn_connection_exampleId]) => aws.ec2transitgateway.getVpnAttachment({
    transitGatewayId: aws_ec2_transit_gateway_exampleId,
    vpnConnectionId: aws_vpn_connection_exampleId,
}, { async: true }));

