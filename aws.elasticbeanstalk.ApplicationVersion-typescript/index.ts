import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultBucket = new aws.s3.Bucket("default", {});
const defaultBucketObject = new aws.s3.BucketObject("default", {
    bucket: defaultBucket.id,
    key: "beanstalk/go-v1.zip",
    source: new pulumi.asset.FileAsset("go-v1.zip"),
});
const defaultApplication = new aws.elasticbeanstalk.Application("default", {
    description: "tf-test-desc",
});
const defaultApplicationVersion = new aws.elasticbeanstalk.ApplicationVersion("default", {
    application: "tf-test-name",
    bucket: defaultBucket.id,
    description: "application version",
    key: defaultBucketObject.id,
});

