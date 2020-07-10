import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.neptune.ParameterGroup("example", {
    family: "neptune1",
    parameters: [{
        name: "neptune_query_timeout",
        value: "25",
    }],
});

