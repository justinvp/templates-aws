import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultDefaultVpcDhcpOptions = new aws.ec2.DefaultVpcDhcpOptions("default", {
    tags: {
        Name: "Default DHCP Option Set",
    },
});

