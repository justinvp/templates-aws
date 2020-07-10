import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleGateway = new aws.directconnect.Gateway("example", {
    amazonSideAsn: "64512",
});
const exampleTransitVirtualInterface = new aws.directconnect.TransitVirtualInterface("example", {
    addressFamily: "ipv4",
    bgpAsn: 65352,
    connectionId: aws_dx_connection_example.id,
    dxGatewayId: exampleGateway.id,
    vlan: 4094,
});

