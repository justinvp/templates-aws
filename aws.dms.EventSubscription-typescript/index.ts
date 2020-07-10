import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.dms.EventSubscription("example", {
    enabled: true,
    eventCategories: [
        "creation",
        "failure",
    ],
    snsTopicArn: aws_sns_topic.example.arn,
    sourceIds: [aws_dms_replication_task.example.replication_task_id],
    sourceType: "replication-task",
    tags: {
        Name: "example",
    },
});

