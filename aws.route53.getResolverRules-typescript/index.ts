import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = pulumi.output(aws.route53.getResolverRules({
    tags: [{
        Environment: "dev",
    }],
}, { async: true }));

