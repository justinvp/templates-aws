import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultCluster = new aws.rds.Cluster("default", {
    availabilityZones: [
        "us-west-2a",
        "us-west-2b",
        "us-west-2c",
    ],
    backupRetentionPeriod: 5,
    clusterIdentifier: "aurora-cluster-demo",
    databaseName: "mydb",
    engine: "aurora-mysql",
    engineVersion: "5.7.mysql_aurora.2.03.2",
    masterPassword: "bar",
    masterUsername: "foo",
    preferredBackupWindow: "07:00-09:00",
});

