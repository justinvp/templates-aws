import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.wafv2.IpSet("example", {
    addresses: [
        "1.2.3.4/32",
        "5.6.7.8/32",
    ],
    description: "Example IP set",
    ipAddressVersion: "IPV4",
    scope: "REGIONAL",
    tags: {
        Tag1: "Value1",
        Tag2: "Value2",
    },
});

