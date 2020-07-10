import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const strict = new aws.iam.AccountPasswordPolicy("strict", {
    allowUsersToChangePassword: true,
    minimumPasswordLength: 8,
    requireLowercaseCharacters: true,
    requireNumbers: true,
    requireSymbols: true,
    requireUppercaseCharacters: true,
});

