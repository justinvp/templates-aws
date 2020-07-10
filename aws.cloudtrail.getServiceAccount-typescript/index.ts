import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = pulumi.output(aws.cloudtrail.getServiceAccount({ async: true }));
const bucket = new aws.s3.Bucket("bucket", {
    forceDestroy: true,
    policy: pulumi.interpolate`{
  "Version": "2008-10-17",
  "Statement": [
    {
      "Sid": "Put bucket policy needed for trails",
      "Effect": "Allow",
      "Principal": {
        "AWS": "${main.arn}"
      },
      "Action": "s3:PutObject",
      "Resource": "arn:aws:s3:::tf-cloudtrail-logging-test-bucket/*"
    },
    {
      "Sid": "Get bucket policy needed for trails",
      "Effect": "Allow",
      "Principal": {
        "AWS": "${main.arn}"
      },
      "Action": "s3:GetBucketAcl",
      "Resource": "arn:aws:s3:::tf-cloudtrail-logging-test-bucket"
    }
  ]
}
`,
});

