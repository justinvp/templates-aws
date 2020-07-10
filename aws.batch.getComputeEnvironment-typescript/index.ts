import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const batch_mongo = pulumi.output(aws.batch.getComputeEnvironment({
    computeEnvironmentName: "batch-mongo-production",
}, { async: true }));

