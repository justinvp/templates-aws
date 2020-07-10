import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cloudtrail.Trail("example", {
    eventSelectors: [{
        dataResources: [{
            type: "AWS::Lambda::Function",
            values: ["arn:aws:lambda"],
        }],
        includeManagementEvents: true,
        readWriteType: "All",
    }],
});

