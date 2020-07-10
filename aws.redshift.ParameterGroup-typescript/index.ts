import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bar = new aws.redshift.ParameterGroup("bar", {
    family: "redshift-1.0",
    parameters: [
        {
            name: "require_ssl",
            value: "true",
        },
        {
            name: "query_group",
            value: "example",
        },
        {
            name: "enable_user_activity_logging",
            value: "true",
        },
    ],
});

