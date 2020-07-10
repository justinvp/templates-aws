import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.servicediscovery.HttpNamespace("example", {
    description: "example",
});

