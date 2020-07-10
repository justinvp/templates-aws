import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const developers = new aws.iam.Group("developers", {
    path: "/users/",
});

