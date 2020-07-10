import pulumi
import pulumi_aws as aws

replicas_target = aws.appautoscaling.Target("replicasTarget",
    max_capacity=15,
    min_capacity=1,
    resource_id=f"cluster:{aws_rds_cluster['example']['id']}",
    scalable_dimension="rds:cluster:ReadReplicaCount",
    service_namespace="rds")
replicas_policy = aws.appautoscaling.Policy("replicasPolicy",
    policy_type="TargetTrackingScaling",
    resource_id=replicas_target.resource_id,
    scalable_dimension=replicas_target.scalable_dimension,
    service_namespace=replicas_target.service_namespace,
    target_tracking_scaling_policy_configuration={
        "predefinedMetricSpecification": {
            "predefinedMetricType": "RDSReaderAverageCPUUtilization",
        },
        "scaleInCooldown": 300,
        "scaleOutCooldown": 300,
        "targetValue": 75,
    })

