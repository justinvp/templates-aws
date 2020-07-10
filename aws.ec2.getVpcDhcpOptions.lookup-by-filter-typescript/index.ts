import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.ec2.getVpcDhcpOptions({
    filters: [
        {
            name: "key",
            values: ["domain-name"],
        },
        {
            name: "value",
            values: ["example.com"],
        },
    ],
}, { async: true }));

