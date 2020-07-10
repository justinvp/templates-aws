import pulumi
import pulumi_aws as aws

default = aws.docdb.Cluster("default",
    availability_zones=[
        "us-west-2a",
        "us-west-2b",
        "us-west-2c",
    ],
    cluster_identifier="docdb-cluster-demo",
    master_password="barbut8chars",
    master_username="foo")
cluster_instances = []
for range in [{"value": i} for i in range(0, 2)]:
    cluster_instances.append(aws.docdb.ClusterInstance(f"clusterInstances-{range['value']}",
        cluster_identifier=default.id,
        identifier=f"docdb-cluster-demo-{range['value']}",
        instance_class="db.r5.large"))

