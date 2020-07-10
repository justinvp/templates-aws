import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.efs.FileSystem("foo", {
    tags: {
        Name: "MyProduct",
    },
});

