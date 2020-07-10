import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultLaunchTemplate = pulumi.output(aws.ec2.getLaunchTemplate({
    name: "my-launch-template",
}, { async: true }));

