import pulumi
import pulumi_aws as aws

example = aws.elasticache.Cluster("example",
    engine="redis",
    engine_version="3.2.10",
    node_type="cache.m4.large",
    num_cache_nodes=1,
    parameter_group_name="default.redis3.2",
    port=6379)

