import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const alias = new aws.iam.AccountAlias("alias", {
    accountAlias: "my-account-alias",
});

