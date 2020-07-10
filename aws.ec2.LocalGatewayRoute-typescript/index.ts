import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.LocalGatewayRoute("example", {
    destinationCidrBlock: "172.16.0.0/16",
    localGatewayRouteTableId: data.aws_ec2_local_gateway_route_table.example.id,
    localGatewayVirtualInterfaceGroupId: data.aws_ec2_local_gateway_virtual_interface_group.example.id,
});

