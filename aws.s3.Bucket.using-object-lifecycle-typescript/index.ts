import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bucket = new aws.s3.Bucket("bucket", {
    acl: "private",
    lifecycleRules: [
        {
            enabled: true,
            expiration: {
                days: 90,
            },
            id: "log",
            prefix: "log/",
            tags: {
                autoclean: "true",
                rule: "log",
            },
            transitions: [
                {
                    days: 30,
                    storageClass: "STANDARD_IA", // or "ONEZONE_IA"
                },
                {
                    days: 60,
                    storageClass: "GLACIER",
                },
            ],
        },
        {
            enabled: true,
            expiration: {
                date: "2016-01-12",
            },
            id: "tmp",
            prefix: "tmp/",
        },
    ],
});
const versioningBucket = new aws.s3.Bucket("versioning_bucket", {
    acl: "private",
    lifecycleRules: [{
        enabled: true,
        noncurrentVersionExpiration: {
            days: 90,
        },
        noncurrentVersionTransitions: [
            {
                days: 30,
                storageClass: "STANDARD_IA",
            },
            {
                days: 60,
                storageClass: "GLACIER",
            },
        ],
        prefix: "config/",
    }],
    versioning: {
        enabled: true,
    },
});

