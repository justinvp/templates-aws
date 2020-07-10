import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultCluster = new aws.rds.Cluster("default", {
    availabilityZones: [
        "us-west-2a",
        "us-west-2b",
        "us-west-2c",
    ],
    clusterIdentifier: "aurora-cluster-demo",
    databaseName: "mydb",
    masterPassword: "barbut8chars",
    masterUsername: "foo",
});
const clusterInstances: aws.rds.ClusterInstance[] = [];
for (let i = 0; i < 2; i++) {
    clusterInstances.push(new aws.rds.ClusterInstance(`cluster_instances-${i}`, {
        clusterIdentifier: defaultCluster.id,
        identifier: `aurora-cluster-demo-${i}`,
        instanceClass: "db.r4.large",
    }));
}

