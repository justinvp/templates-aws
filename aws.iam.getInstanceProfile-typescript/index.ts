import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.iam.getInstanceProfile({
    name: "an_example_instance_profile_name",
}, { async: true }));

