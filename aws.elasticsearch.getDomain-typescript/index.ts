import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myDomain = pulumi.output(aws.elasticsearch.getDomain({
    domainName: "my-domain-name",
}, { async: true }));

