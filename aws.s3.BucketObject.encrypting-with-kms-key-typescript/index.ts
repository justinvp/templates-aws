import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const examplekms = new aws.kms.Key("examplekms", {
    deletionWindowInDays: 7,
    description: "KMS key 1",
});
const examplebucket = new aws.s3.Bucket("examplebucket", {
    acl: "private",
});
const examplebucketObject = new aws.s3.BucketObject("examplebucket_object", {
    bucket: examplebucket.id,
    key: "someobject",
    kmsKeyId: examplekms.arn,
    source: new pulumi.asset.FileAsset("index.html"),
});

