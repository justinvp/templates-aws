import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.ec2.CustomerGateway("main", {
    bgpAsn: 65000,
    ipAddress: "172.83.124.10",
    tags: {
        Name: "main-customer-gateway",
    },
    type: "ipsec.1",
});

