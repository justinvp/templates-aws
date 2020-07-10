import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.dax.ParameterGroup("example", {
    parameters: [
        {
            name: "query-ttl-millis",
            value: "100000",
        },
        {
            name: "record-ttl-millis",
            value: "100000",
        },
    ],
});

