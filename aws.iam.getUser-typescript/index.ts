import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.iam.getUser({
    userName: "an_example_user_name",
}, { async: true }));

