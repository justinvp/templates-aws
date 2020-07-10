import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleGateway = new aws.directconnect.Gateway("example", {
    amazonSideAsn: "64512",
});
const exampleVpc = new aws.ec2.Vpc("example", {
    cidrBlock: "10.255.255.0/28",
});
const exampleVpnGateway = new aws.ec2.VpnGateway("example", {
    vpcId: exampleVpc.id,
});
const exampleGatewayAssociation = new aws.directconnect.GatewayAssociation("example", {
    allowedPrefixes: [
        "210.52.109.0/24",
        "175.45.176.0/22",
    ],
    associatedGatewayId: exampleVpnGateway.id,
    dxGatewayId: exampleGateway.id,
});

