import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const mainRouteTableAssociation = new aws.ec2.MainRouteTableAssociation("a", {
    routeTableId: aws_route_table_bar.id,
    vpcId: aws_vpc_foo.id,
});

