import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const cache = new aws.opsworks.MemcachedLayer("cache", {
    stackId: aws_opsworks_stack_main.id,
});

