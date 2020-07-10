import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultBucket = new aws.s3.Bucket("default", {});
const defaultSpotDatafeedSubscription = new aws.ec2.SpotDatafeedSubscription("default", {
    bucket: defaultBucket.bucket,
    prefix: "my_subdirectory",
});

