import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ebs.Volume("example", {
    availabilityZone: "us-west-2a",
    size: 40,
    tags: {
        Name: "HelloWorld",
    },
});
const exampleSnapshot = new aws.ebs.Snapshot("example_snapshot", {
    tags: {
        Name: "HelloWorld_snap",
    },
    volumeId: example.id,
});
const exampleCopy = new aws.ebs.SnapshotCopy("example_copy", {
    sourceRegion: "us-west-2",
    sourceSnapshotId: exampleSnapshot.id,
    tags: {
        Name: "HelloWorld_copy_snap",
    },
});

