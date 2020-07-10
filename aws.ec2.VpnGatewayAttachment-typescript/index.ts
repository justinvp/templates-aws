import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const network = new aws.ec2.Vpc("network", {
    cidrBlock: "10.0.0.0/16",
});
const vpn = new aws.ec2.VpnGateway("vpn", {
    tags: {
        Name: "example-vpn-gateway",
    },
});
const vpnAttachment = new aws.ec2.VpnGatewayAttachment("vpn_attachment", {
    vpcId: network.id,
    vpnGatewayId: vpn.id,
});

