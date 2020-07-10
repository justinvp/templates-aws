import pulumi
import pulumi_aws as aws

example = aws.elasticache.Cluster("example",
    engine="memcached",
    node_type="cache.m4.large",
    num_cache_nodes=2,
    parameter_group_name="default.memcached1.4",
    port=11211)

