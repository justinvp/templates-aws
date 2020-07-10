import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.iam.getGroup({
    groupName: "an_example_group_name",
}, { async: true }));

