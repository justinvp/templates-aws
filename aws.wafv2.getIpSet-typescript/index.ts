import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.wafv2.getIpSet({
    name: "some-ip-set",
    scope: "REGIONAL",
}, { async: true }));

