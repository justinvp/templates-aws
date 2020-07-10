import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const important_bucket = pulumi.output(aws.s3.getBucket({
    bucket: "important-bucket",
}, { async: true }));
const example = new aws.cloudtrail.Trail("example", {
    eventSelectors: [{
        dataResources: [{
            type: "AWS::S3::Object",
            // Make sure to append a trailing '/' to your ARN if you want
            // to monitor all objects in a bucket.
            values: [pulumi.interpolate`${important_bucket.arn}/`],
        }],
        includeManagementEvents: true,
        readWriteType: "All",
    }],
});

