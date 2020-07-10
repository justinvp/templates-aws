import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.backup.Selection("example", {
    iamRoleArn: aws_iam_role_example.arn,
    planId: aws_backup_plan_example.id,
    selectionTags: [{
        key: "foo",
        type: "STRINGEQUALS",
        value: "bar",
    }],
});

