import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const accepter = new aws.Provider("accepter", {});
const accepterCallerIdentity = pulumi.output(aws.getCallerIdentity({ provider: accepter, async: true }));
// Accepter's side of the VIF.
const example = new aws.directconnect.Gateway("example", {
    amazonSideAsn: "64512",
}, { provider: accepter });
// Creator's side of the VIF
const creator = new aws.directconnect.HostedTransitVirtualInterface("creator", {
    addressFamily: "ipv4",
    bgpAsn: 65352,
    connectionId: "dxcon-zzzzzzzz",
    ownerAccountId: accepterCallerIdentity.accountId,
    vlan: 4094,
}, { dependsOn: [example] });
const accepterHostedTransitVirtualInterfaceAcceptor = new aws.directconnect.HostedTransitVirtualInterfaceAcceptor("accepter", {
    dxGatewayId: example.id,
    tags: {
        Side: "Accepter",
    },
    virtualInterfaceId: creator.id,
}, { provider: accepter });

