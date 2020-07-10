import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultCluster = new aws.docdb.Cluster("default", {
    availabilityZones: [
        "us-west-2a",
        "us-west-2b",
        "us-west-2c",
    ],
    clusterIdentifier: "docdb-cluster-demo",
    masterPassword: "barbut8chars",
    masterUsername: "foo",
});
const clusterInstances: aws.docdb.ClusterInstance[] = [];
for (let i = 0; i < 2; i++) {
    clusterInstances.push(new aws.docdb.ClusterInstance(`cluster_instances-${i}`, {
        clusterIdentifier: defaultCluster.id,
        identifier: `docdb-cluster-demo-${i}`,
        instanceClass: "db.r5.large",
    }));
}

