import pulumi
import pulumi_aws as aws

replicas = aws.appautoscaling.Target("replicas",
    max_capacity=15,
    min_capacity=1,
    resource_id=f"cluster:{aws_rds_cluster['example']['id']}",
    scalable_dimension="rds:cluster:ReadReplicaCount",
    service_namespace="rds")

