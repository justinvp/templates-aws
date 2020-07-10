import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const replicasTarget = new aws.appautoscaling.Target("replicas", {
    maxCapacity: 15,
    minCapacity: 1,
    resourceId: pulumi.interpolate`cluster:${aws_rds_cluster_example.id}`,
    scalableDimension: "rds:cluster:ReadReplicaCount",
    serviceNamespace: "rds",
});
const replicasPolicy = new aws.appautoscaling.Policy("replicas", {
    policyType: "TargetTrackingScaling",
    resourceId: replicasTarget.resourceId,
    scalableDimension: replicasTarget.scalableDimension,
    serviceNamespace: replicasTarget.serviceNamespace,
    targetTrackingScalingPolicyConfiguration: {
        predefinedMetricSpecification: {
            predefinedMetricType: "RDSReaderAverageCPUUtilization",
        },
        scaleInCooldown: 300,
        scaleOutCooldown: 300,
        targetValue: 75,
    },
});

