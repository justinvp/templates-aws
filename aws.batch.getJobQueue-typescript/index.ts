import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test_queue = pulumi.output(aws.batch.getJobQueue({
    name: "tf-test-batch-job-queue",
}, { async: true }));

