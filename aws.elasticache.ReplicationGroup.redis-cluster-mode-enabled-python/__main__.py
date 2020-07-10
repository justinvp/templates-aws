import pulumi
import pulumi_aws as aws

baz = aws.elasticache.ReplicationGroup("baz",
    automatic_failover_enabled=True,
    cluster_mode={
        "numNodeGroups": 2,
        "replicasPerNodeGroup": 1,
    },
    node_type="cache.t2.small",
    parameter_group_name="default.redis3.2.cluster.on",
    port=6379,
    replication_group_description="test description")

