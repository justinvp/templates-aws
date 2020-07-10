import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.workspaces.getBundle({
    name: "Value with Windows 10 and Office 2016",
    owner: "AMAZON",
}, { async: true }));

