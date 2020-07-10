import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const routeTable = new aws.ec2.RouteTable("r", {
    routes: [
        {
            cidrBlock: "10.0.1.0/24",
            gatewayId: aws_internet_gateway_main.id,
        },
        {
            egressOnlyGatewayId: aws_egress_only_internet_gateway_foo.id,
            ipv6CidrBlock: "::/0",
        },
    ],
    tags: {
        Name: "main",
    },
    vpcId: aws_vpc_default.id,
});

