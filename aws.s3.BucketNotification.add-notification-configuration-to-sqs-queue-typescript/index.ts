import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bucket = new aws.s3.Bucket("bucket", {});
const queue = new aws.sqs.Queue("queue", {
    policy: pulumi.interpolate`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": "*",
      "Action": "sqs:SendMessage",
	  "Resource": "arn:aws:sqs:*:*:s3-event-notification-queue",
      "Condition": {
        "ArnEquals": { "aws:SourceArn": "${bucket.arn}" }
      }
    }
  ]
}
`,
});
const bucketNotification = new aws.s3.BucketNotification("bucket_notification", {
    bucket: bucket.id,
    queues: [{
        events: ["s3:ObjectCreated:*"],
        filterSuffix: ".log",
        queueArn: queue.arn,
    }],
});

