import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myStackPermission = new aws.opsworks.Permission("my_stack_permission", {
    allowSsh: true,
    allowSudo: true,
    level: "iam_only",
    stackId: aws_opsworks_stack_stack.id,
    userArn: aws_iam_user_user.arn,
});

