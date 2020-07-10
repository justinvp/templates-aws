import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.ec2transitgateway.getTransitGateway({
    filters: [{
        name: "options.amazon-side-asn",
        values: ["64512"],
    }],
}, { async: true }));

