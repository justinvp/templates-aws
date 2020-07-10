import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testQueue = new aws.batch.JobQueue("test_queue", {
    computeEnvironments: [
        aws_batch_compute_environment_test_environment_1.arn,
        aws_batch_compute_environment_test_environment_2.arn,
    ],
    priority: 1,
    state: "ENABLED",
});

