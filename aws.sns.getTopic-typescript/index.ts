import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.sns.getTopic({
    name: "an_example_topic",
}, { async: true }));

