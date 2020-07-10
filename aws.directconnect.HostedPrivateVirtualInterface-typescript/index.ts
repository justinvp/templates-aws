import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.directconnect.HostedPrivateVirtualInterface("foo", {
    addressFamily: "ipv4",
    bgpAsn: 65352,
    connectionId: "dxcon-zzzzzzzz",
    vlan: 4094,
});

