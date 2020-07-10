import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const lambda_example = new aws.lb.TargetGroup("lambda-example", {
    targetType: "lambda",
});

