import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const logBucket = new aws.s3.Bucket("log_bucket", {
    acl: "log-delivery-write",
});
const bucket = new aws.s3.Bucket("b", {
    acl: "private",
    loggings: [{
        targetBucket: logBucket.id,
        targetPrefix: "log/",
    }],
});

