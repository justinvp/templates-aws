import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.servicequotas.getService({
    serviceName: "Amazon Virtual Private Cloud (Amazon VPC)",
}, { async: true }));

