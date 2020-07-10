import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = pulumi.output(aws.ec2.getInstance({
    filters: [
        {
            name: "image-id",
            values: ["ami-xxxxxxxx"],
        },
        {
            name: "tag:Name",
            values: ["instance-name-tag"],
        },
    ],
    instanceId: "i-instanceid",
}, { async: true }));

