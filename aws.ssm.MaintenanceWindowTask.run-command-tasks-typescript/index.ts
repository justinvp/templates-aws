import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ssm.MaintenanceWindowTask("example", {
    maxConcurrency: "2",
    maxErrors: "1",
    priority: 1,
    serviceRoleArn: aws_iam_role_example.arn,
    targets: [{
        key: "InstanceIds",
        values: [aws_instance_example.id],
    }],
    taskArn: "AWS-RunShellScript",
    taskInvocationParameters: {
        runCommandParameters: {
            notificationConfig: {
                notificationArn: aws_sns_topic_example.arn,
                notificationEvents: ["All"],
                notificationType: "Command",
            },
            outputS3Bucket: aws_s3_bucket_example.bucket,
            outputS3KeyPrefix: "output",
            parameters: [{
                name: "commands",
                values: ["date"],
            }],
            serviceRoleArn: aws_iam_role_example.arn,
            timeoutSeconds: 600,
        },
    },
    taskType: "RUN_COMMAND",
    windowId: aws_ssm_maintenance_window_example.id,
});

