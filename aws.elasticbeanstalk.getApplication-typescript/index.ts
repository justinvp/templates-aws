import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.elasticbeanstalk.getApplication({
    name: "example",
}, { async: true }));

export const arn = example.arn;
export const description = example.description;

