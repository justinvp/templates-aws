import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultParameterGroup = new aws.elasticache.ParameterGroup("default", {
    family: "redis2.8",
    parameters: [
        {
            name: "activerehashing",
            value: "yes",
        },
        {
            name: "min-slaves-to-write",
            value: "2",
        },
    ],
});

