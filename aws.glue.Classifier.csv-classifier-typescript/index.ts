import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Classifier("example", {
    csvClassifier: {
        allowSingleColumn: false,
        containsHeader: "PRESENT",
        delimiter: ",",
        disableValueTrimming: false,
        headers: [
            "example1",
            "example2",
        ],
        quoteSymbol: "'",
    },
});

