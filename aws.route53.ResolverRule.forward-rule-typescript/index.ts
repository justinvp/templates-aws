import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const fwd = new aws.route53.ResolverRule("fwd", {
    domainName: "example.com",
    resolverEndpointId: aws_route53_resolver_endpoint_foo.id,
    ruleType: "FORWARD",
    tags: {
        Environment: "Prod",
    },
    targetIps: [{
        ip: "123.45.67.89",
    }],
});

