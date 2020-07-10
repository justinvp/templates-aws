import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.all([aws_dx_gateway_example.id, aws_ec2_transit_gateway_example.id]).apply(([aws_dx_gateway_exampleId, aws_ec2_transit_gateway_exampleId]) => aws.ec2transitgateway.getDirectConnectGatewayAttachment({
    dxGatewayId: aws_dx_gateway_exampleId,
    transitGatewayId: aws_ec2_transit_gateway_exampleId,
}, { async: true }));

