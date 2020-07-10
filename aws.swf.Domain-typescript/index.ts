import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.swf.Domain("foo", {
    description: "SWF Domain",
    workflowExecutionRetentionPeriodInDays: "30",
});

