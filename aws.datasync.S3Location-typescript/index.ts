import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.datasync.S3Location("example", {
    s3BucketArn: aws_s3_bucket_example.arn,
    s3Config: {
        bucketAccessRoleArn: aws_iam_role_example.arn,
    },
    subdirectory: "/example/prefix",
});

