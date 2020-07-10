import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleGateway = new aws.directconnect.Gateway("example", {
    amazonSideAsn: "64512",
});
const exampleTransitGateway = new aws.ec2transitgateway.TransitGateway("example", {});
const exampleGatewayAssociation = new aws.directconnect.GatewayAssociation("example", {
    allowedPrefixes: [
        "10.255.255.0/30",
        "10.255.255.8/30",
    ],
    associatedGatewayId: exampleTransitGateway.id,
    dxGatewayId: exampleGateway.id,
});

