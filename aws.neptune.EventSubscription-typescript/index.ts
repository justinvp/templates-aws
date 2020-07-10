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
const example = new aws.neptune.ClusterInstance("example", {
    applyImmediately: true,
    clusterIdentifier: defaultCluster.id,
    engine: "neptune",
    instanceClass: "db.r4.large",
});
const defaultTopic = new aws.sns.Topic("default", {});
const defaultEventSubscription = new aws.neptune.EventSubscription("default", {
    eventCategories: [
        "maintenance",
        "availability",
        "creation",
        "backup",
        "restoration",
        "recovery",
        "deletion",
        "failover",
        "failure",
        "notification",
        "configuration change",
        "read replica",
    ],
    snsTopicArn: defaultTopic.arn,
    sourceIds: [example.id],
    sourceType: "db-instance",
    tags: {
        env: "test",
    },
});

