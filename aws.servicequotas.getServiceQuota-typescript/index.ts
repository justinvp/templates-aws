import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const byQuotaCode = pulumi.output(aws.servicequotas.getServiceQuota({
    quotaCode: "L-F678F1CE",
    serviceCode: "vpc",
}, { async: true }));
const byQuotaName = pulumi.output(aws.servicequotas.getServiceQuota({
    quotaName: "VPCs per Region",
    serviceCode: "vpc",
}, { async: true }));

