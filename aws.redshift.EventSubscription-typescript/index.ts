import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultCluster = new aws.redshift.Cluster("default", {
    clusterIdentifier: "default",
    databaseName: "default",
});
const defaultTopic = new aws.sns.Topic("default", {});
const defaultEventSubscription = new aws.redshift.EventSubscription("default", {
    eventCategories: [
        "configuration",
        "management",
        "monitoring",
        "security",
    ],
    severity: "INFO",
    snsTopicArn: defaultTopic.arn,
    sourceIds: [defaultCluster.id],
    sourceType: "cluster",
    tags: {
        Name: "default",
    },
});

