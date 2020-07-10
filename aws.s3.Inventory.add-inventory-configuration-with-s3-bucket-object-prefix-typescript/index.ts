import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.s3.Bucket("test", {});
const inventory = new aws.s3.Bucket("inventory", {});
const test_prefix = new aws.s3.Inventory("test-prefix", {
    bucket: test.id,
    destination: {
        bucket: {
            bucketArn: inventory.arn,
            format: "ORC",
            prefix: "inventory",
        },
    },
    filter: {
        prefix: "documents/",
    },
    includedObjectVersions: "All",
    schedule: {
        frequency: "Daily",
    },
});

