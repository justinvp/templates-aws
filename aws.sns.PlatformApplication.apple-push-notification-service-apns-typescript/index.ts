import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const apnsApplication = new aws.sns.PlatformApplication("apns_application", {
    platform: "APNS",
    platformCredential: "<APNS PRIVATE KEY>",
    platformPrincipal: "<APNS CERTIFICATE>",
});

