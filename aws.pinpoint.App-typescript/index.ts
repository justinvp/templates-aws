import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.pinpoint.App("example", {
    limits: {
        maximumDuration: 600,
    },
    quietTime: {
        end: "06:00",
        start: "00:00",
    },
});

