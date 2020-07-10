import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.ec2.Vpc("foo", {
    cidrBlock: "10.0.0.0/16",
});
const alphaSubnet = new aws.ec2.Subnet("alpha", {
    availabilityZone: "us-west-2a",
    cidrBlock: "10.0.1.0/24",
    vpcId: foo.id,
});
const alphaMountTarget = new aws.efs.MountTarget("alpha", {
    fileSystemId: aws_efs_file_system_foo.id,
    subnetId: alphaSubnet.id,
});

