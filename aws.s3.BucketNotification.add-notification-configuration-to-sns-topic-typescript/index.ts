import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bucket = new aws.s3.Bucket("bucket", {});
const topic = new aws.sns.Topic("topic", {
    policy: pulumi.interpolate`{
    "Version":"2012-10-17",
    "Statement":[{
        "Effect": "Allow",
        "Principal": {"AWS":"*"},
        "Action": "SNS:Publish",
        "Resource": "arn:aws:sns:*:*:s3-event-notification-topic",
        "Condition":{
            "ArnLike":{"aws:SourceArn":"${bucket.arn}"}
        }
    }]
}
`,
});
const bucketNotification = new aws.s3.BucketNotification("bucket_notification", {
    bucket: bucket.id,
    topics: [{
        events: ["s3:ObjectCreated:*"],
        filterSuffix: ".log",
        topicArn: topic.arn,
    }],
});

