import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.directconnect.HostedTransitVirtualInterface("example", {
    addressFamily: "ipv4",
    bgpAsn: 65352,
    connectionId: aws_dx_connection_example.id,
    vlan: 4094,
});

