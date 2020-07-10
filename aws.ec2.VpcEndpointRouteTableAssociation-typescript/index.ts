import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.VpcEndpointRouteTableAssociation("example", {
    routeTableId: aws_route_table_example.id,
    vpcEndpointId: aws_vpc_endpoint_example.id,
});

