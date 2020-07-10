import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bar = new aws.inspector.ResourceGroup("bar", {
    tags: {
        Env: "bar",
        Name: "foo",
    },
});
const foo = new aws.inspector.AssessmentTarget("foo", {
    resourceGroupArn: bar.arn,
});

