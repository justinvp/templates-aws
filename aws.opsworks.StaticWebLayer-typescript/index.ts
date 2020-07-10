import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const web = new aws.opsworks.StaticWebLayer("web", {
    stackId: aws_opsworks_stack_main.id,
});

