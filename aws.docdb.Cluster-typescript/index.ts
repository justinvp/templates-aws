import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const docdb = new aws.docdb.Cluster("docdb", {
    backupRetentionPeriod: 5,
    clusterIdentifier: "my-docdb-cluster",
    engine: "docdb",
    masterPassword: "mustbeeightchars",
    masterUsername: "foo",
    preferredBackupWindow: "07:00-09:00",
    skipFinalSnapshot: true,
});

