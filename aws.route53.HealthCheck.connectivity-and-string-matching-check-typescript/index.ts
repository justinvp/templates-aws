import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.route53.HealthCheck("example", {
    failureThreshold: 5,
    fqdn: "example.com",
    port: 443,
    requestInterval: 30,
    resourcePath: "/",
    searchString: "example",
    type: "HTTPS_STR_MATCH",
});

