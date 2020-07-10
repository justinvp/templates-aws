import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.ec2transitgateway.getRouteTable({
    filters: [
        {
            name: "default-association-route-table",
            values: ["true"],
        },
        {
            name: "transit-gateway-id",
            values: ["tgw-12345678"],
        },
    ],
}, { async: true }));

