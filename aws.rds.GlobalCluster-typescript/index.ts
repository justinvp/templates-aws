import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const primary = new aws.Provider("primary", {
    region: "us-east-2",
});
const secondary = new aws.Provider("secondary", {
    region: "us-west-2",
});
const example = new aws.rds.GlobalCluster("example", {
    globalClusterIdentifier: "example",
}, { provider: primary });
const primaryCluster = new aws.rds.Cluster("primary", {
    // ... other configuration ...
    engineMode: "global",
    globalClusterIdentifier: example.id,
}, { provider: primary });
const primaryClusterInstance = new aws.rds.ClusterInstance("primary", {
    // ... other configuration ...
    clusterIdentifier: primaryCluster.id,
}, { provider: primary });
const secondaryCluster = new aws.rds.Cluster("secondary", {
    // ... other configuration ...
    engineMode: "global",
    globalClusterIdentifier: example.id,
}, { provider: secondary, dependsOn: [primaryClusterInstance] });
const secondaryClusterInstance = new aws.rds.ClusterInstance("secondary", {
    // ... other configuration ...
    clusterIdentifier: secondaryCluster.id,
}, { provider: secondary });

