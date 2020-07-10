import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Trigger("example", {
    actions: [{
        jobName: aws_glue_job_example1.name,
    }],
    predicate: {
        conditions: [{
            crawlState: "SUCCEEDED",
            crawlerName: aws_glue_crawler_example2.name,
        }],
    },
    type: "CONDITIONAL",
});

