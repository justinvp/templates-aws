import pulumi
import pulumi_aws as aws

default = aws.rds.Cluster("default",
    availability_zones=[
        "us-west-2a",
        "us-west-2b",
        "us-west-2c",
    ],
    cluster_identifier="aurora-cluster-demo",
    database_name="mydb",
    master_password="barbut8chars",
    master_username="foo")
cluster_instances = []
for range in [{"value": i} for i in range(0, 2)]:
    cluster_instances.append(aws.rds.ClusterInstance(f"clusterInstances-{range['value']}",
        cluster_identifier=default.id,
        identifier=f"aurora-cluster-demo-{range['value']}",
        instance_class="db.r4.large"))

