import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const postgresql = new aws.rds.Cluster("postgresql", {
    availabilityZones: [
        "us-west-2a",
        "us-west-2b",
        "us-west-2c",
    ],
    backupRetentionPeriod: 5,
    clusterIdentifier: "aurora-cluster-demo",
    databaseName: "mydb",
    engine: "aurora-postgresql",
    masterPassword: "bar",
    masterUsername: "foo",
    preferredBackupWindow: "07:00-09:00",
});

