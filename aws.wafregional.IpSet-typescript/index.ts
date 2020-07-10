import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ipset = new aws.wafregional.IpSet("ipset", {
    ipSetDescriptors: [
        {
            type: "IPV4",
            value: "192.0.7.0/24",
        },
        {
            type: "IPV4",
            value: "10.16.16.0/16",
        },
    ],
});

