import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const replicas = new aws.appautoscaling.Target("replicas", {
    maxCapacity: 15,
    minCapacity: 1,
    resourceId: pulumi.interpolate`cluster:${aws_rds_cluster_example.id}`,
    scalableDimension: "rds:cluster:ReadReplicaCount",
    serviceNamespace: "rds",
});

