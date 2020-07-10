import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.backup.Selection("example", {
    iamRoleArn: aws_iam_role_example.arn,
    planId: aws_backup_plan_example.id,
    resources: [
        aws_db_instance_example.arn,
        aws_ebs_volume_example.arn,
        aws_efs_file_system_example.arn,
    ],
});

