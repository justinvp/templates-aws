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
    masterPassword: "bar",
    masterUsername: "foo",
    preferredBackupWindow: "07:00-09:00",
});
const test1 = new aws.rds.ClusterInstance("test1", {
    applyImmediately: true,
    clusterIdentifier: defaultCluster.id,
    identifier: "test1",
    instanceClass: "db.t2.small",
});
const test2 = new aws.rds.ClusterInstance("test2", {
    applyImmediately: true,
    clusterIdentifier: defaultCluster.id,
    identifier: "test2",
    instanceClass: "db.t2.small",
});
const test3 = new aws.rds.ClusterInstance("test3", {
    applyImmediately: true,
    clusterIdentifier: defaultCluster.id,
    identifier: "test3",
    instanceClass: "db.t2.small",
});
const eligible = new aws.rds.ClusterEndpoint("eligible", {
    clusterEndpointIdentifier: "reader",
    clusterIdentifier: defaultCluster.id,
    customEndpointType: "READER",
    excludedMembers: [
        test1.id,
        test2.id,
    ],
});
const static = new aws.rds.ClusterEndpoint("static", {
    clusterEndpointIdentifier: "static",
    clusterIdentifier: defaultCluster.id,
    customEndpointType: "READER",
    staticMembers: [
        test1.id,
        test3.id,
    ],
});

