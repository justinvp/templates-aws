import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const mainVpc = new aws.ec2.Vpc("main", {
    cidrBlock: "10.0.0.0/16",
});
const private_a = new aws.ec2.Subnet("private-a", {
    availabilityZone: "us-east-1a",
    cidrBlock: "10.0.0.0/24",
    vpcId: mainVpc.id,
});
const private_b = new aws.ec2.Subnet("private-b", {
    availabilityZone: "us-east-1b",
    cidrBlock: "10.0.1.0/24",
    vpcId: mainVpc.id,
});
const mainDirectory = new aws.directoryservice.Directory("main", {
    password: "#S1ncerely",
    size: "Small",
    vpcSettings: {
        subnetIds: [
            private_a.id,
            private_b.id,
        ],
        vpcId: mainVpc.id,
    },
});
const mainWorkspacesDirectory = new aws.workspaces.Directory("main", {
    directoryId: mainDirectory.id,
    selfServicePermissions: {
        increaseVolumeSize: true,
        rebuildWorkspace: true,
    },
});

