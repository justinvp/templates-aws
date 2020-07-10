import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const domainTest = new aws.lightsail.Domain("domain_test", {
    domainName: "mydomain.com",
});

