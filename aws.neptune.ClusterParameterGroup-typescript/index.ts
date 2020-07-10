import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.neptune.ClusterParameterGroup("example", {
    description: "neptune cluster parameter group",
    family: "neptune1",
    parameters: [{
        name: "neptune_enable_audit_log",
        value: "1",
    }],
});

