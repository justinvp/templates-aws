import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultRouteTable = new aws.ec2.DefaultRouteTable("r", {
    defaultRouteTableId: aws_vpc_foo.defaultRouteTableId,
    routes: [{}],
    tags: {
        Name: "default table",
    },
});

