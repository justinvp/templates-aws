import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const examplebucket = new aws.s3.Bucket("examplebucket", {
    acl: "private",
    objectLockConfiguration: {
        objectLockEnabled: "Enabled",
    },
    versioning: {
        enabled: true,
    },
});
const examplebucketObject = new aws.s3.BucketObject("examplebucket_object", {
    bucket: examplebucket.id,
    forceDestroy: true,
    key: "someobject",
    objectLockLegalHoldStatus: "ON",
    objectLockMode: "GOVERNANCE",
    objectLockRetainUntilDate: "2021-12-31T23:59:60Z",
    source: new pulumi.asset.FileAsset("important.txt"),
});

