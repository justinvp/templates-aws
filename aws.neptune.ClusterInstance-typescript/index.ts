import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultCluster = new aws.neptune.Cluster("default", {
    applyImmediately: true,
    backupRetentionPeriod: 5,
    clusterIdentifier: "neptune-cluster-demo",
    engine: "neptune",
    iamDatabaseAuthenticationEnabled: true,
    preferredBackupWindow: "07:00-09:00",
    skipFinalSnapshot: true,
});
const example: aws.neptune.ClusterInstance[] = [];
for (let i = 0; i < 2; i++) {
    example.push(new aws.neptune.ClusterInstance(`example-${i}`, {
        applyImmediately: true,
        clusterIdentifier: defaultCluster.id,
        engine: "neptune",
        instanceClass: "db.r4.large",
    }));
}

