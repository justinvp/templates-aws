import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.rds.ClusterSnapshot("example", {
    dbClusterIdentifier: aws_rds_cluster_example.id,
    dbClusterSnapshotIdentifier: "resourcetestsnapshot1234",
});

