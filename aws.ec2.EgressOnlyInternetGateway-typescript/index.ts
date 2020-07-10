import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleVpc = new aws.ec2.Vpc("example", {
    assignGeneratedIpv6CidrBlock: true,
    cidrBlock: "10.1.0.0/16",
});
const exampleEgressOnlyInternetGateway = new aws.ec2.EgressOnlyInternetGateway("example", {
    tags: {
        Name: "main",
    },
    vpcId: exampleVpc.id,
});

