import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const account = new aws.organizations.Account("account", {
    email: "john@doe.org",
});

