import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const app = new aws.opsworks.PhpAppLayer("app", {
    stackId: aws_opsworks_stack_main.id,
});

