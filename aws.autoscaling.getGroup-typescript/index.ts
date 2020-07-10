import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = pulumi.output(aws.autoscaling.getGroup({
    name: "foo",
}, { async: true }));

