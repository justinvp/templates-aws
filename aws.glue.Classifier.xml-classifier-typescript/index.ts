import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Classifier("example", {
    xmlClassifier: {
        classification: "example",
        rowTag: "example",
    },
});

