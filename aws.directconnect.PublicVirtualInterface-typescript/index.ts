import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.directconnect.PublicVirtualInterface("foo", {
    addressFamily: "ipv4",
    amazonAddress: "175.45.176.2/30",
    bgpAsn: 65352,
    connectionId: "dxcon-zzzzzzzz",
    customerAddress: "175.45.176.1/30",
    routeFilterPrefixes: [
        "210.52.109.0/24",
        "175.45.176.0/22",
    ],
    vlan: 4094,
});

