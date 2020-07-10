import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = pulumi.output(aws.getBillingServiceAccount({ async: true }));
const billingLogs = new aws.s3.Bucket("billing_logs", {
    acl: "private",
    policy: pulumi.interpolate`{
  "Id": "Policy",
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "s3:GetBucketAcl", "s3:GetBucketPolicy"
      ],
      "Effect": "Allow",
      "Resource": "arn:aws:s3:::my-billing-tf-test-bucket",
      "Principal": {
        "AWS": [
          "${main.arn}"
        ]
      }
    },
    {
      "Action": [
        "s3:PutObject"
      ],
      "Effect": "Allow",
      "Resource": "arn:aws:s3:::my-billing-tf-test-bucket/*",
      "Principal": {
        "AWS": [
          "${main.arn}"
        ]
      }
    }
  ]
}
`,
});

