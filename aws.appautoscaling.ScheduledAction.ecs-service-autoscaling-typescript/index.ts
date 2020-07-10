import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ecsTarget = new aws.appautoscaling.Target("ecs", {
    maxCapacity: 4,
    minCapacity: 1,
    resourceId: "service/clusterName/serviceName",
    scalableDimension: "ecs:service:DesiredCount",
    serviceNamespace: "ecs",
});
const ecsScheduledAction = new aws.appautoscaling.ScheduledAction("ecs", {
    resourceId: ecsTarget.resourceId,
    scalableDimension: ecsTarget.scalableDimension,
    scalableTargetAction: {
        maxCapacity: 10,
        minCapacity: 1,
    },
    schedule: "at(2006-01-02T15:04:05)",
    serviceNamespace: ecsTarget.serviceNamespace,
});

