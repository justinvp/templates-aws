import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const lb = new aws.opsworks.HaproxyLayer("lb", {
    stackId: aws_opsworks_stack_main.id,
    statsPassword: "foobarbaz",
});

