import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const db = new aws.opsworks.MysqlLayer("db", {
    stackId: aws_opsworks_stack_main.id,
});

