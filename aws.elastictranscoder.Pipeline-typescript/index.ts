import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bar = new aws.elastictranscoder.Pipeline("bar", {
    contentConfig: {
        bucket: aws_s3_bucket_content_bucket.bucket,
        storageClass: "Standard",
    },
    inputBucket: aws_s3_bucket_input_bucket.bucket,
    role: aws_iam_role_test_role.arn,
    thumbnailConfig: {
        bucket: aws_s3_bucket_thumb_bucket.bucket,
        storageClass: "Standard",
    },
});

