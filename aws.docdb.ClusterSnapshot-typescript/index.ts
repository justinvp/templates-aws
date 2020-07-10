import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.docdb.ClusterSnapshot("example", {
    dbClusterIdentifier: aws_docdb_cluster_example.id,
    dbClusterSnapshotIdentifier: "resourcetestsnapshot1234",
});

