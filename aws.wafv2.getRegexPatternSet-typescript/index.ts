import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.wafv2.getRegexPatternSet({
    name: "some-regex-pattern-set",
    scope: "REGIONAL",
}, { async: true }));

