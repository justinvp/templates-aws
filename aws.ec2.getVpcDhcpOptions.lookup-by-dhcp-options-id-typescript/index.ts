import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.ec2.getVpcDhcpOptions({
    dhcpOptionsId: "dopts-12345678",
}, { async: true }));

