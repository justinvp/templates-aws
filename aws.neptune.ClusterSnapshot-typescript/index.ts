import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.neptune.ClusterSnapshot("example", {
    dbClusterIdentifier: aws_neptune_cluster_example.id,
    dbClusterSnapshotIdentifier: "resourcetestsnapshot1234",
});

