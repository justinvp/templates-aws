import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const fooLocalGateways = pulumi.output(aws.ec2.getLocalGateways({
    tags: {
        service: "production",
    },
}, { async: true }));

export const foo = fooLocalGateways.ids;

