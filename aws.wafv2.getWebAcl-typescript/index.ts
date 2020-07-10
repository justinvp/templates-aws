import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.wafv2.getWebAcl({
    name: "some-web-acl",
    scope: "REGIONAL",
}, { async: true }));

