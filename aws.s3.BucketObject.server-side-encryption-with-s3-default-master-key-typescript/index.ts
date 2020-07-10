import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const examplebucket = new aws.s3.Bucket("examplebucket", {
    acl: "private",
});
const examplebucketObject = new aws.s3.BucketObject("examplebucket_object", {
    bucket: examplebucket.id,
    key: "someobject",
    serverSideEncryption: "aws:kms",
    source: new pulumi.asset.FileAsset("index.html"),
});

