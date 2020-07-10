import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = pulumi.output(aws.ssm.getDocument({
    documentFormat: "YAML",
    name: "AWS-GatherSoftwareInventory",
}, { async: true }));

export const content = foo.content;

