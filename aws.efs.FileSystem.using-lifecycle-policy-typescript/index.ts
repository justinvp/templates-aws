import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const fooWithLifecylePolicy = new aws.efs.FileSystem("foo_with_lifecyle_policy", {
    lifecyclePolicy: {
        transitionToIa: "AFTER_30_DAYS",
    },
});

