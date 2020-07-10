import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Declare the data source
const pc = aws_vpc_foo.id.apply(id => aws.ec2.getVpcPeeringConnection({
    peerCidrBlock: "10.0.1.0/22",
    vpcId: id,
}, { async: true }));
// Create a route table
const rt = new aws.ec2.RouteTable("rt", {
    vpcId: aws_vpc_foo.id,
});
// Create a route
const route = new aws.ec2.Route("r", {
    destinationCidrBlock: pc.peerCidrBlock!,
    routeTableId: rt.id,
    vpcPeeringConnectionId: pc.id!,
});

