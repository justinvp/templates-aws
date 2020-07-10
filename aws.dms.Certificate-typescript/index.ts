import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create a new certificate
const test = new aws.dms.Certificate("test", {
    certificateId: "test-dms-certificate-tf",
    certificatePem: "...",
});

