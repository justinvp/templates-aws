import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const local = new aws.Provider("local", {region: "us-east-1"});
const peer = new aws.Provider("peer", {region: "us-west-2"});
const peerRegion = aws.getRegion({});
const localTransitGateway = new aws.ec2transitgateway.TransitGateway("localTransitGateway", {tags: {
    Name: "Local TGW",
}}, {
    provider: aws.local,
});
const peerTransitGateway = new aws.ec2transitgateway.TransitGateway("peerTransitGateway", {tags: {
    Name: "Peer TGW",
}}, {
    provider: aws.peer,
});
const example = new aws.ec2transitgateway.PeeringAttachment("example", {
    peerAccountId: peerTransitGateway.ownerId,
    peerRegion: peerRegion.then(peerRegion => peerRegion.name),
    peerTransitGatewayId: peerTransitGateway.id,
    transitGatewayId: localTransitGateway.id,
    tags: {
        Name: "TGW Peering Requestor",
    },
});

