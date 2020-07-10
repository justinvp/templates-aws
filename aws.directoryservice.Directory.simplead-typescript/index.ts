import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.ec2.Vpc("main", {
    cidrBlock: "10.0.0.0/16",
});
const foo = new aws.ec2.Subnet("foo", {
    availabilityZone: "us-west-2a",
    cidrBlock: "10.0.1.0/24",
    vpcId: main.id,
});
const barSubnet = new aws.ec2.Subnet("bar", {
    availabilityZone: "us-west-2b",
    cidrBlock: "10.0.2.0/24",
    vpcId: main.id,
});
const barDirectory = new aws.directoryservice.Directory("bar", {
    password: "SuperSecretPassw0rd",
    size: "Small",
    tags: {
        Project: "foo",
    },
    vpcSettings: {
        subnetIds: [
            foo.id,
            barSubnet.id,
        ],
        vpcId: main.id,
    },
});

