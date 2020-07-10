import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const app = new aws.pinpoint.App("app", {});
const testStream = new aws.kinesis.Stream("test_stream", {
    shardCount: 1,
});
const testRole = new aws.iam.Role("test_role", {
    assumeRolePolicy: `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "pinpoint.us-east-1.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
`,
});
const stream = new aws.pinpoint.EventStream("stream", {
    applicationId: app.applicationId,
    destinationStreamArn: testStream.arn,
    roleArn: testRole.arn,
});
const testRolePolicy = new aws.iam.RolePolicy("test_role_policy", {
    policy: `{
  "Version": "2012-10-17",
  "Statement": {
    "Action": [
      "kinesis:PutRecords",
      "kinesis:DescribeStream"
    ],
    "Effect": "Allow",
    "Resource": [
      "arn:aws:kinesis:us-east-1:*:*/*"
    ]
  }
}
`,
    role: testRole.id,
});

