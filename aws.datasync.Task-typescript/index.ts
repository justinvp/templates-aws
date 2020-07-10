import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.datasync.Task("example", {
    destinationLocationArn: aws_datasync_location_s3_destination.arn,
    options: {
        bytesPerSecond: -1,
    },
    sourceLocationArn: aws_datasync_location_nfs_source.arn,
});

