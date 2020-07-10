import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const app = new aws.pinpoint.App("app", {});
const channel = new aws.pinpoint.AdmChannel("channel", {
    applicationId: app.applicationId,
    clientId: "",
    clientSecret: "",
    enabled: true,
});

