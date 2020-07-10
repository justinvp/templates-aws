import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const routeTableAssociation = new aws.ec2.RouteTableAssociation("routeTableAssociation", {
    subnetId: aws_subnet.foo.id,
    routeTableId: aws_route_table.bar.id,
});

