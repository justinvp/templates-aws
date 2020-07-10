import pulumi
import pulumi_aws as aws

ecs_target = aws.appautoscaling.Target("ecsTarget",
    max_capacity=4,
    min_capacity=1,
    resource_id="service/clusterName/serviceName",
    scalable_dimension="ecs:service:DesiredCount",
    service_namespace="ecs")
ecs_policy = aws.appautoscaling.Policy("ecsPolicy",
    policy_type="StepScaling",
    resource_id=ecs_target.resource_id,
    scalable_dimension=ecs_target.scalable_dimension,
    service_namespace=ecs_target.service_namespace,
    step_scaling_policy_configuration={
        "adjustment_type": "ChangeInCapacity",
        "cooldown": 60,
        "metric_aggregation_type": "Maximum",
        "stepAdjustment": [{
            "metricIntervalUpperBound": 0,
            "scaling_adjustment": -1,
        }],
    })

