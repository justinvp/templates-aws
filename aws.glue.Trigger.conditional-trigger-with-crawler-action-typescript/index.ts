import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Trigger("example", {
    actions: [{
        crawlerName: aws_glue_crawler_example1.name,
    }],
    predicate: {
        conditions: [{
            jobName: aws_glue_job_example2.name,
            state: "SUCCEEDED",
        }],
    },
    type: "CONDITIONAL",
});

