import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Classifier("example", {
    grokClassifier: {
        classification: "example",
        grokPattern: "example",
    },
});

