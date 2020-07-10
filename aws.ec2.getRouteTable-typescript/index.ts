import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const subnetId = config.require("subnetId");

const selected = pulumi.output(aws.ec2.getRouteTable({
    subnetId: subnetId,
}, { async: true }));
const route = new aws.ec2.Route("route", {
    destinationCidrBlock: "10.0.1.0/22",
    routeTableId: selected.id,
    vpcPeeringConnectionId: "pcx-45ff3dc1",
});

