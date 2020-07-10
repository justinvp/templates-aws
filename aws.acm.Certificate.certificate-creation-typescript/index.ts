import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const cert = new aws.acm.Certificate("cert", {
    domainName: "example.com",
    tags: {
        Environment: "test",
    },
    validationMethod: "DNS",
});

