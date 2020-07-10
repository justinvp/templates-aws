import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.inspector.ResourceGroup("example", {
    tags: {
        Env: "bar",
        Name: "foo",
    },
});

