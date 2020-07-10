import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const route = new aws.ec2.Route("route", {
    routeTableId: "rtb-4fbb3ac4",
    destinationCidrBlock: "10.0.1.0/22",
    vpcPeeringConnectionId: "pcx-45ff3dc1",
}, {
    dependsOn: ["aws_route_table.testing"],
});

