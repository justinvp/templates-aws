import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const yada = new aws.cloudwatch.LogGroup("yada", {
    tags: {
        Application: "serviceA",
        Environment: "production",
    },
});

