import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testSnapshotCopyGrant = new aws.redshift.SnapshotCopyGrant("test", {
    snapshotCopyGrantName: "my-grant",
});
const testCluster = new aws.redshift.Cluster("test", {
    // ... other configuration ...
    snapshotCopy: {
        destinationRegion: "us-east-2",
        grantName: testSnapshotCopyGrant.snapshotCopyGrantName,
    },
});

