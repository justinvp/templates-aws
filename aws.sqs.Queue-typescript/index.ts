import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const queue = new aws.sqs.Queue("queue", {
    delaySeconds: 90,
    maxMessageSize: 2048,
    messageRetentionSeconds: 86400,
    receiveWaitTimeSeconds: 10,
    redrivePolicy: JSON.stringify({
        deadLetterTargetArn: aws_sqs_queue.queue_deadletter.arn,
        maxReceiveCount: 4,
    }),
    tags: {
        Environment: "production",
    },
});

