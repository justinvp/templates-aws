import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apigateway.RestApi("example", {
    endpointConfiguration: {
        types: "REGIONAL",
    },
});

