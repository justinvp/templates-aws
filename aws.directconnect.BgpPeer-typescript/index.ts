import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const peer = new aws.directconnect.BgpPeer("peer", {
    addressFamily: "ipv6",
    bgpAsn: 65351,
    virtualInterfaceId: aws_dx_private_virtual_interface_foo.id,
});

