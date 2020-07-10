import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ebs.Volume("example", {
    availabilityZone: "us-west-2a",
    size: 40,
});
const exampleSnapshot = new aws.ebs.Snapshot("example_snapshot", {
    volumeId: example.id,
});
const examplePerm = new aws.ec2.SnapshotCreateVolumePermission("example_perm", {
    accountId: "12345678",
    snapshotId: exampleSnapshot.id,
});

