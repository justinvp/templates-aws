import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bucket = new aws.s3.Bucket("bucket", {
    acl: "private",
});
const firehoseRole = new aws.iam.Role("firehose_role", {
    assumeRolePolicy: `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "firehose.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
`,
});
const testStream = new aws.kinesis.FirehoseDeliveryStream("test_stream", {
    destination: "s3",
    s3Configuration: {
        bucketArn: bucket.arn,
        roleArn: firehoseRole.arn,
    },
});

