import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const master = new aws.guardduty.Detector("master", {
    enable: true,
});
const bucket = new aws.s3.Bucket("bucket", {
    acl: "private",
});
const myIPSetBucketObject = new aws.s3.BucketObject("MyIPSet", {
    acl: "public-read",
    bucket: bucket.id,
    content: "10.0.0.0/8\n",
    key: "MyIPSet",
});
const myIPSetIPSet = new aws.guardduty.IPSet("MyIPSet", {
    activate: true,
    detectorId: master.id,
    format: "TXT",
    location: pulumi.interpolate`https://s3.amazonaws.com/${myIPSetBucketObject.bucket}/${myIPSetBucketObject.key}`,
});

