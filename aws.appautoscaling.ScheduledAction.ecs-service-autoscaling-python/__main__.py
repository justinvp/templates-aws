import pulumi
import pulumi_aws as aws

ecs_target = aws.appautoscaling.Target("ecsTarget",
    max_capacity=4,
    min_capacity=1,
    resource_id="service/clusterName/serviceName",
    scalable_dimension="ecs:service:DesiredCount",
    service_namespace="ecs")
ecs_scheduled_action = aws.appautoscaling.ScheduledAction("ecsScheduledAction",
    resource_id=ecs_target.resource_id,
    scalable_dimension=ecs_target.scalable_dimension,
    scalable_target_action={
        "max_capacity": 10,
        "min_capacity": 1,
    },
    schedule="at(2006-01-02T15:04:05)",
    service_namespace=ecs_target.service_namespace)

