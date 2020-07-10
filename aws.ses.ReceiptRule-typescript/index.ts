import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Add a header to the email and store it in S3
const store = new aws.ses.ReceiptRule("store", {
    addHeaderActions: [{
        headerName: "Custom-Header",
        headerValue: "Added by SES",
        position: 1,
    }],
    enabled: true,
    recipients: ["karen@example.com"],
    ruleSetName: "default-rule-set",
    s3Actions: [{
        bucketName: "emails",
        position: 2,
    }],
    scanEnabled: true,
});

