import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Workflow("example", {});
const example_start = new aws.glue.Trigger("example-start", {
    actions: [{
        jobName: "example-job",
    }],
    type: "ON_DEMAND",
    workflowName: example.name,
});
const example_inner = new aws.glue.Trigger("example-inner", {
    actions: [{
        jobName: "another-example-job",
    }],
    predicate: {
        conditions: [{
            jobName: "example-job",
            state: "SUCCEEDED",
        }],
    },
    type: "CONDITIONAL",
    workflowName: example.name,
});

