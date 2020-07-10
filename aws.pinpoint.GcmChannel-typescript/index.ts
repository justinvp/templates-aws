import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const app = new aws.pinpoint.App("app", {});
const gcm = new aws.pinpoint.GcmChannel("gcm", {
    apiKey: "api_key",
    applicationId: app.applicationId,
});

