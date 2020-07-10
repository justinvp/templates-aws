import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const monitor = new aws.opsworks.GangliaLayer("monitor", {
    password: "foobarbaz",
    stackId: aws_opsworks_stack_main.id,
});

