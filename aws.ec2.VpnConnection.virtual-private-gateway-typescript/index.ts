import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const vpc = new aws.ec2.Vpc("vpc", {
    cidrBlock: "10.0.0.0/16",
});
const vpnGateway = new aws.ec2.VpnGateway("vpn_gateway", {
    vpcId: vpc.id,
});
const customerGateway = new aws.ec2.CustomerGateway("customer_gateway", {
    bgpAsn: 65000,
    ipAddress: "172.0.0.1",
    type: "ipsec.1",
});
const main = new aws.ec2.VpnConnection("main", {
    customerGatewayId: customerGateway.id,
    staticRoutesOnly: true,
    type: "ipsec.1",
    vpnGatewayId: vpnGateway.id,
});

