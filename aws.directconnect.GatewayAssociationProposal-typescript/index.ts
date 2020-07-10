import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.directconnect.GatewayAssociationProposal("example", {
    associatedGatewayId: aws_vpn_gateway_example.id,
    dxGatewayId: aws_dx_gateway_example.id,
    dxGatewayOwnerAccountId: aws_dx_gateway_example.ownerAccountId,
});

