import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const production = new aws.ssm.MaintenanceWindow("production", {
    cutoff: 1,
    duration: 3,
    schedule: "cron(0 16 ? * TUE *)",
});

