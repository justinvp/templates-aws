import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const org = new aws.organizations.Organization("org", {
    awsServiceAccessPrincipals: [
        "cloudtrail.amazonaws.com",
        "config.amazonaws.com",
    ],
    featureSet: "ALL",
});

