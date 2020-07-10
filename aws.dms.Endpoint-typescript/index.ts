import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create a new endpoint
const test = new aws.dms.Endpoint("test", {
    certificateArn: "arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012",
    databaseName: "test",
    endpointId: "test-dms-endpoint-tf",
    endpointType: "source",
    engineName: "aurora",
    extraConnectionAttributes: "",
    kmsKeyArn: "arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012",
    password: "test",
    port: 3306,
    serverName: "test",
    sslMode: "none",
    tags: {
        Name: "test",
    },
    username: "test",
});

