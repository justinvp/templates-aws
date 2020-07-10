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
    taskArn: "AWS-RestartEC2Instance",
    taskInvocationParameters: {
        automationParameters: {
            documentVersion: "$LATEST",
            parameters: [{
                name: "InstanceId",
                values: [aws_instance_example.id],
            }],
        },
    },
    taskType: "AUTOMATION",
    windowId: aws_ssm_maintenance_window_example.id,
});

