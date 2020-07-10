import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bucket = new aws.s3.Bucket("b", {
    acl: "public-read",
    corsRules: [{
        allowedHeaders: ["*"],
        allowedMethods: [
            "PUT",
            "POST",
        ],
        allowedOrigins: ["https://s3-website-test.mydomain.com"],
        exposeHeaders: ["ETag"],
        maxAgeSeconds: 3000,
    }],
});

