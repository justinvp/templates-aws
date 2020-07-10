import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.macie.S3BucketAssociation("example", {
    bucketName: "tf-macie-example",
    classificationType: {
        oneTime: "FULL",
    },
    prefix: "data",
});

