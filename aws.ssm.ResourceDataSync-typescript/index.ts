import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const hogeBucket = new aws.s3.Bucket("hoge", {
    region: "us-east-1",
});
const hogeBucketPolicy = new aws.s3.BucketPolicy("hoge", {
    bucket: hogeBucket.bucket,
    policy: `{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "SSMBucketPermissionsCheck",
            "Effect": "Allow",
            "Principal": {
                "Service": "ssm.amazonaws.com"
            },
            "Action": "s3:GetBucketAcl",
            "Resource": "arn:aws:s3:::tf-test-bucket-1234"
        },
        {
            "Sid": " SSMBucketDelivery",
            "Effect": "Allow",
            "Principal": {
                "Service": "ssm.amazonaws.com"
            },
            "Action": "s3:PutObject",
            "Resource": ["arn:aws:s3:::tf-test-bucket-1234/*"],
            "Condition": {
                "StringEquals": {
                    "s3:x-amz-acl": "bucket-owner-full-control"
                }
            }
        }
    ]
}
`,
});
const foo = new aws.ssm.ResourceDataSync("foo", {
    s3Destination: {
        bucketName: hogeBucket.bucket,
        region: hogeBucket.region,
    },
});

