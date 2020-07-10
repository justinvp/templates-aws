import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const master = new aws.guardduty.Detector("master", {
    enable: true,
});
const bucket = new aws.s3.Bucket("bucket", {
    acl: "private",
});
const myThreatIntelSetBucketObject = new aws.s3.BucketObject("MyThreatIntelSet", {
    acl: "public-read",
    bucket: bucket.id,
    content: "10.0.0.0/8\n",
    key: "MyThreatIntelSet",
});
const myThreatIntelSetThreatIntelSet = new aws.guardduty.ThreatIntelSet("MyThreatIntelSet", {
    activate: true,
    detectorId: master.id,
    format: "TXT",
    location: pulumi.interpolate`https://s3.amazonaws.com/${myThreatIntelSetBucketObject.bucket}/${myThreatIntelSetBucketObject.key}`,
});

