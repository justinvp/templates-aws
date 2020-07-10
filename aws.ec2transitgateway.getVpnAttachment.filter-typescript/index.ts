import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = pulumi.output(aws.ec2transitgateway.getVpnAttachment({
    filters: [{
        name: "resource-id",
        values: ["some-resource"],
    }],
}, { async: true }));

