import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const peer = new aws.Provider("peer", {
    region: "us-west-2",
});
const main = new aws.ec2.Vpc("main", {
    cidrBlock: "10.0.0.0/16",
});
const peerVpc = new aws.ec2.Vpc("peer", {
    cidrBlock: "10.1.0.0/16",
}, { provider: peer });
const peerCallerIdentity = pulumi.output(aws.getCallerIdentity({ provider: peer, async: true }));
// Requester's side of the connection.
const peerVpcPeeringConnection = new aws.ec2.VpcPeeringConnection("peer", {
    autoAccept: false,
    peerOwnerId: peerCallerIdentity.accountId,
    peerRegion: "us-west-2",
    peerVpcId: peerVpc.id,
    tags: {
        Side: "Requester",
    },
    vpcId: main.id,
});
// Accepter's side of the connection.
const peerVpcPeeringConnectionAccepter = new aws.ec2.VpcPeeringConnectionAccepter("peer", {
    autoAccept: true,
    tags: {
        Side: "Accepter",
    },
    vpcPeeringConnectionId: peerVpcPeeringConnection.id,
}, { provider: peer });

