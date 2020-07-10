import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const accepter = new aws.Provider("accepter", {});
const accepterCallerIdentity = pulumi.output(aws.getCallerIdentity({ provider: accepter, async: true }));
// Accepter's side of the VIF.
const vpnGw = new aws.ec2.VpnGateway("vpn_gw", {}, { provider: accepter });
// Creator's side of the VIF
const creator = new aws.directconnect.HostedPrivateVirtualInterface("creator", {
    addressFamily: "ipv4",
    bgpAsn: 65352,
    connectionId: "dxcon-zzzzzzzz",
    ownerAccountId: accepterCallerIdentity.accountId,
    vlan: 4094,
}, { dependsOn: [vpnGw] });
const accepterHostedPrivateVirtualInterfaceAccepter = new aws.directconnect.HostedPrivateVirtualInterfaceAccepter("accepter", {
    tags: {
        Side: "Accepter",
    },
    virtualInterfaceId: creator.id,
    vpnGatewayId: vpnGw.id,
}, { provider: accepter });

