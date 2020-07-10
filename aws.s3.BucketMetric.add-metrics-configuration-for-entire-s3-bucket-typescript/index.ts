import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.s3.Bucket("example", {});
const example_entire_bucket = new aws.s3.BucketMetric("example-entire-bucket", {
    bucket: example.bucket,
});

