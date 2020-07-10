import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.ram.getResourceShare({
    name: "example",
    resourceOwner: "SELF",
}, { async: true }));

