import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const fooVpc = new aws.ec2.Vpc("foo", {
    cidrBlock: "10.1.0.0/16",
});
const fooSubnet = new aws.ec2.Subnet("foo", {
    availabilityZone: "us-west-2a",
    cidrBlock: "10.1.1.0/24",
    tags: {
        Name: "tf-dbsubnet-test-1",
    },
    vpcId: fooVpc.id,
});
const bar = new aws.ec2.Subnet("bar", {
    availabilityZone: "us-west-2b",
    cidrBlock: "10.1.2.0/24",
    tags: {
        Name: "tf-dbsubnet-test-2",
    },
    vpcId: fooVpc.id,
});
const fooSubnetGroup = new aws.redshift.SubnetGroup("foo", {
    subnetIds: [
        fooSubnet.id,
        bar.id,
    ],
    tags: {
        environment: "Production",
    },
});

