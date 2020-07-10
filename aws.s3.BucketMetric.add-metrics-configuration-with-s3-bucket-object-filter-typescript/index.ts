import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.s3.Bucket("example", {});
const example_filtered = new aws.s3.BucketMetric("example-filtered", {
    bucket: example.bucket,
    filter: {
        prefix: "documents/",
        tags: {
            class: "blue",
            priority: "high",
        },
    },
});

