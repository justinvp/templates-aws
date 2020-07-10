import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.iam.getRole({
    name: "an_example_role_name",
}, { async: true }));

