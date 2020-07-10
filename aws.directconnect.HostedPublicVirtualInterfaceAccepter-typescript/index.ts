import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const accepter = new aws.Provider("accepter", {});
const accepterCallerIdentity = pulumi.output(aws.getCallerIdentity({ provider: accepter, async: true }));
// Creator's side of the VIF
const creator = new aws.directconnect.HostedPublicVirtualInterface("creator", {
    addressFamily: "ipv4",
    amazonAddress: "175.45.176.2/30",
    bgpAsn: 65352,
    connectionId: "dxcon-zzzzzzzz",
    customerAddress: "175.45.176.1/30",
    ownerAccountId: accepterCallerIdentity.accountId,
    routeFilterPrefixes: [
        "210.52.109.0/24",
        "175.45.176.0/22",
    ],
    vlan: 4094,
});
// Accepter's side of the VIF.
const accepterHostedPublicVirtualInterfaceAccepter = new aws.directconnect.HostedPublicVirtualInterfaceAccepter("accepter", {
    tags: {
        Side: "Accepter",
    },
    virtualInterfaceId: creator.id,
}, { provider: accepter });

