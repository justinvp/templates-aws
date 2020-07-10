import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.waf.getRateBasedRule({
    name: "tfWAFRateBasedRule",
}, { async: true }));

