import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = pulumi.output(aws.getPrefixList({
    filters: [{
        name: "prefix-list-id",
        values: ["pl-68a54001"],
    }],
}, { async: true }));

