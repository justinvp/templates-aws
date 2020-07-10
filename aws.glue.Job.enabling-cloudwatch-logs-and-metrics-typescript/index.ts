import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleLogGroup = new aws.cloudwatch.LogGroup("example", {
    retentionInDays: 14,
});
const exampleJob = new aws.glue.Job("example", {
    defaultArguments: {
        // ... potentially other arguments ...
        "--continuous-log-logGroup": exampleLogGroup.name,
        "--enable-continuous-cloudwatch-log": "true",
        "--enable-continuous-log-filter": "true",
        "--enable-metrics": "",
    },
});

