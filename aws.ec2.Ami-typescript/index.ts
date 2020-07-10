import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create an AMI that will start a machine whose root device is backed by
// an EBS volume populated from a snapshot. It is assumed that such a snapshot
// already exists with the id "snap-xxxxxxxx".
const example = new aws.ec2.Ami("example", {
    ebsBlockDevices: [{
        deviceName: "/dev/xvda",
        snapshotId: "snap-xxxxxxxx",
        volumeSize: 8,
    }],
    rootDeviceName: "/dev/xvda",
    virtualizationType: "hvm",
});

