import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.wafregional.getIpset({
    name: "tfWAFRegionalIPSet",
}, { async: true }));

