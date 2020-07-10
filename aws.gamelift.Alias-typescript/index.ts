import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.gamelift.Alias("example", {
    description: "Example Description",
    routingStrategy: {
        message: "Example Message",
        type: "TERMINAL",
    },
});

