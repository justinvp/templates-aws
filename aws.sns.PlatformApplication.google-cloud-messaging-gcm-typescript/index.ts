import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const gcmApplication = new aws.sns.PlatformApplication("gcm_application", {
    platform: "GCM",
    platformCredential: "<GCM API KEY>",
});

