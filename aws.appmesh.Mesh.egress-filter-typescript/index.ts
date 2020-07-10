import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const simple = new aws.appmesh.Mesh("simple", {
    spec: {
        egressFilter: {
            type: "ALLOW_ALL",
        },
    },
});

