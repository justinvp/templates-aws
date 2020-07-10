import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ses.EmailIdentity("example", {
    email: "email@example.com",
});

