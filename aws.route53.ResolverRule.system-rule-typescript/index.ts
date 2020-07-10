import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const sys = new aws.route53.ResolverRule("sys", {
    domainName: "subdomain.example.com",
    ruleType: "SYSTEM",
});

