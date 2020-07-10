import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.iot.Thing("example", {
    attributes: {
        First: "examplevalue",
    },
});

