import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.route53.getResolverRule({
    domainName: "subdomain.example.com",
    ruleType: "SYSTEM",
}, { async: true }));

