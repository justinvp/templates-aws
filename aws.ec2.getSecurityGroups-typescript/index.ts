import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = pulumi.output(aws.ec2.getSecurityGroups({
    tags: {
        Application: "k8s",
        Environment: "dev",
    },
}, { async: true }));

