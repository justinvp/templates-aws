import pulumi
import pulumi_aws as aws

ecs_target = aws.appautoscaling.Target("ecsTarget",
    max_capacity=4,
    min_capacity=1,
    resource_id=f"service/{aws_ecs_cluster['example']['name']}/{aws_ecs_service['example']['name']}",
    scalable_dimension="ecs:service:DesiredCount",
    service_namespace="ecs")

