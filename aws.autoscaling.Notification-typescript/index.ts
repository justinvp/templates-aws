import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.sns.Topic("example", {});
const bar = new aws.autoscaling.Group("bar", {});
const foo = new aws.autoscaling.Group("foo", {});
const exampleNotifications = new aws.autoscaling.Notification("example_notifications", {
    groupNames: [
        bar.name,
        foo.name,
    ],
    notifications: [
        "autoscaling:EC2_INSTANCE_LAUNCH",
        "autoscaling:EC2_INSTANCE_TERMINATE",
        "autoscaling:EC2_INSTANCE_LAUNCH_ERROR",
        "autoscaling:EC2_INSTANCE_TERMINATE_ERROR",
    ],
    topicArn: example.arn,
});

