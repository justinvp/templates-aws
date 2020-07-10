import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const dada = new aws.cloudwatch.LogGroup("dada", {});
const yada = new aws.cloudwatch.LogMetricFilter("yada", {
    logGroupName: dada.name,
    metricTransformation: {
        name: "EventCount",
        namespace: "YourNamespace",
        value: "1",
    },
    pattern: "",
});

