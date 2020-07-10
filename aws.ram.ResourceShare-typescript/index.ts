import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ram.ResourceShare("example", {
    allowExternalPrincipals: true,
    tags: {
        Environment: "Production",
    },
});

