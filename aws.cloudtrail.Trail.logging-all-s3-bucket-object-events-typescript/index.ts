import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cloudtrail.Trail("example", {
    eventSelectors: [{
        dataResources: [{
            type: "AWS::S3::Object",
            values: ["arn:aws:s3:::"],
        }],
        includeManagementEvents: true,
        readWriteType: "All",
    }],
});

