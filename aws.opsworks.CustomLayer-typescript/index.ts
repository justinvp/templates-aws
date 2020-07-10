import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const custlayer = new aws.opsworks.CustomLayer("custlayer", {
    shortName: "awesome",
    stackId: aws_opsworks_stack_main.id,
});

