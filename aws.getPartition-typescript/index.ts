import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = pulumi.output(aws.getPartition({ async: true }));
const s3Policy = current.apply(current => aws.iam.getPolicyDocument({
    statements: [{
        actions: ["s3:ListBucket"],
        resources: [`arn:${current.partition}:s3:::my-bucket`],
        sid: "1",
    }],
}, { async: true }));

