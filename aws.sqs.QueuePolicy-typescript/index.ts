import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const queue = new aws.sqs.Queue("q", {});
const test = new aws.sqs.QueuePolicy("test", {
    policy: pulumi.interpolate`{
  "Version": "2012-10-17",
  "Id": "sqspolicy",
  "Statement": [
    {
      "Sid": "First",
      "Effect": "Allow",
      "Principal": "*",
      "Action": "sqs:SendMessage",
      "Resource": "${queue.arn}",
      "Condition": {
        "ArnEquals": {
          "aws:SourceArn": "${aws_sns_topic_example.arn}"
        }
      }
    }
  ]
}
`,
    queueUrl: queue.id,
});

