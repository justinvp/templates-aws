import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleVpc = new aws.ec2.Vpc("example", {
    cidrBlock: "10.0.0.0/16",
});
const examplePrivateDnsNamespace = new aws.servicediscovery.PrivateDnsNamespace("example", {
    description: "example",
    vpc: exampleVpc.id,
});

