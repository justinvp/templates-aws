import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testBucket = new aws.s3.Bucket("test", {});
const inventory = new aws.s3.Bucket("inventory", {});
const testInventory = new aws.s3.Inventory("test", {
    bucket: testBucket.id,
    destination: {
        bucket: {
            bucketArn: inventory.arn,
            format: "ORC",
        },
    },
    includedObjectVersions: "All",
    schedule: {
        frequency: "Daily",
    },
});

