import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const clusterName = config.get("clusterName") || "example";

const exampleLogGroup = new aws.cloudwatch.LogGroup("example", {
    retentionInDays: 7,
});
const exampleCluster = new aws.eks.Cluster("example", {
    enabledClusterLogTypes: [
        "api",
        "audit",
    ],
}, { dependsOn: [exampleLogGroup] });

