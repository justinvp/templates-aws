import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const subnetId = config.require("subnetId");

const selected = pulumi.output(aws.ec2.getRouteTable({
    subnetId: subnetId,
}, { async: true }));
const route = aws_route_table_selected.id.apply(id => aws.ec2.getRoute({
    destinationCidrBlock: "10.0.1.0/24",
    routeTableId: id,
}, { async: true }));
const interfaceNetworkInterface = route.apply(route => aws.ec2.getNetworkInterface({
    networkInterfaceId: route.networkInterfaceId!,
}, { async: true }));

