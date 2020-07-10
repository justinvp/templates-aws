import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const dnsResolver = new aws.ec2.VpcDhcpOptions("dns_resolver", {
    domainNameServers: [
        "8.8.8.8",
        "8.8.4.4",
    ],
});

