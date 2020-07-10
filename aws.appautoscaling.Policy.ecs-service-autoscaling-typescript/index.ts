import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ecsTarget = new aws.appautoscaling.Target("ecs_target", {
    maxCapacity: 4,
    minCapacity: 1,
    resourceId: "service/clusterName/serviceName",
    scalableDimension: "ecs:service:DesiredCount",
    serviceNamespace: "ecs",
});
const ecsPolicy = new aws.appautoscaling.Policy("ecs_policy", {
    policyType: "StepScaling",
    resourceId: ecsTarget.resourceId,
    scalableDimension: ecsTarget.scalableDimension,
    serviceNamespace: ecsTarget.serviceNamespace,
    stepScalingPolicyConfiguration: {
        adjustmentType: "ChangeInCapacity",
        cooldown: 60,
        metricAggregationType: "Maximum",
        stepAdjustments: [{
            metricIntervalUpperBound: "0",
            scalingAdjustment: -1,
        }],
    },
});

