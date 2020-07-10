import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const app = new aws.pinpoint.App("app", {});
const sms = new aws.pinpoint.SmsChannel("sms", {
    applicationId: app.applicationId,
});

