import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ecsTarget = new aws.appautoscaling.Target("ecs_target", {
    maxCapacity: 4,
    minCapacity: 1,
    resourceId: pulumi.interpolate`service/${aws_ecs_cluster_example.name}/${aws_ecs_service_example.name}`,
    scalableDimension: "ecs:service:DesiredCount",
    serviceNamespace: "ecs",
});

