import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.servicequotas.ServiceQuota("example", {
    quotaCode: "L-F678F1CE",
    serviceCode: "vpc",
    value: 75,
});

