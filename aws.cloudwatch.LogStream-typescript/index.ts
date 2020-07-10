import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const yada = new aws.cloudwatch.LogGroup("yada", {});
const foo = new aws.cloudwatch.LogStream("foo", {
    logGroupName: yada.name,
});

