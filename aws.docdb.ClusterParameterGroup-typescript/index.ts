import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.docdb.ClusterParameterGroup("example", {
    description: "docdb cluster parameter group",
    family: "docdb3.6",
    parameters: [{
        name: "tls",
        value: "enabled",
    }],
});

