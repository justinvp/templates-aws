import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.ec2.VpcPeeringConnection("foo", {
    peerOwnerId: var_peer_owner_id,
    peerVpcId: aws_vpc_bar.id,
    vpcId: aws_vpc_foo.id,
});

