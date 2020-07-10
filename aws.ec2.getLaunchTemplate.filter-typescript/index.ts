import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = pulumi.output(aws.ec2.getLaunchTemplate({
    filters: [{
        name: "launch-template-name",
        values: ["some-template"],
    }],
}, { async: true }));

