import pulumi
import pulumi_aws as aws

default = aws.neptune.Cluster("default",
    apply_immediately=True,
    backup_retention_period=5,
    cluster_identifier="neptune-cluster-demo",
    engine="neptune",
    iam_database_authentication_enabled=True,
    preferred_backup_window="07:00-09:00",
    skip_final_snapshot=True)
example = []
for range in [{"value": i} for i in range(0, 2)]:
    example.append(aws.neptune.ClusterInstance(f"example-{range['value']}",
        apply_immediately=True,
        cluster_identifier=default.id,
        engine="neptune",
        instance_class="db.r4.large"))

