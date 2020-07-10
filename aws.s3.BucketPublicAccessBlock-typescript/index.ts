import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleBucket = new aws.s3.Bucket("example", {});
const exampleBucketPublicAccessBlock = new aws.s3.BucketPublicAccessBlock("example", {
    blockPublicAcls: true,
    blockPublicPolicy: true,
    bucket: exampleBucket.id,
});

