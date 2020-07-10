import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const fooVpc = new aws.ec2.Vpc("foo", {
    cidrBlock: "10.0.0.0/16",
    tags: {
        Name: "tf-test",
    },
});
const fooSubnet = new aws.ec2.Subnet("foo", {
    availabilityZone: "us-west-2a",
    cidrBlock: "10.0.0.0/24",
    tags: {
        Name: "tf-test",
    },
    vpcId: fooVpc.id,
});
const bar = new aws.elasticache.SubnetGroup("bar", {
    subnetIds: [fooSubnet.id],
});

