import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ubuntu = pulumi.output(aws.ec2.getLaunchConfiguration({
    name: "test-launch-config",
}, { async: true }));

