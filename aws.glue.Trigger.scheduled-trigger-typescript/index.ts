import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Trigger("example", {
    actions: [{
        jobName: aws_glue_job_example.name,
    }],
    schedule: "cron(15 12 * * ? *)",
    type: "SCHEDULED",
});

