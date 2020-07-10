import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const vpnGw = new aws.ec2.VpnGateway("vpn_gw", {
    tags: {
        Name: "main",
    },
    vpcId: aws_vpc_main.id,
});

