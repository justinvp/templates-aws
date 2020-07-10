import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = pulumi.output(aws.ec2.getVpcEndpointService({
    filters: [{
        name: "service-name",
        values: ["some-service"],
    }],
}, { async: true }));

