import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleVpc = new aws.ec2.Vpc("example", {
    cidrBlock: "10.0.0.0/16",
    enableDnsHostnames: true,
    enableDnsSupport: true,
});
const examplePrivateDnsNamespace = new aws.servicediscovery.PrivateDnsNamespace("example", {
    description: "example",
    vpc: exampleVpc.id,
});
const exampleService = new aws.servicediscovery.Service("example", {
    dnsConfig: {
        dnsRecords: [{
            ttl: 10,
            type: "A",
        }],
        namespaceId: examplePrivateDnsNamespace.id,
        routingPolicy: "MULTIVALUE",
    },
    healthCheckCustomConfig: {
        failureThreshold: 1,
    },
});

