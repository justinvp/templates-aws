import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const groups = pulumi.output(aws.getAutoscalingGroups({
    filters: [
        {
            name: "key",
            values: ["Team"],
        },
        {
            name: "value",
            values: ["Pets"],
        },
    ],
}, { async: true }));
const slackNotifications = new aws.autoscaling.Notification("slack_notifications", {
    groupNames: groups.names,
    notifications: [
        "autoscaling:EC2_INSTANCE_LAUNCH",
        "autoscaling:EC2_INSTANCE_TERMINATE",
        "autoscaling:EC2_INSTANCE_LAUNCH_ERROR",
        "autoscaling:EC2_INSTANCE_TERMINATE_ERROR",
    ],
    topicArn: "TOPIC ARN",
});

