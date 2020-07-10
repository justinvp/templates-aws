import pulumi
import pulumi_aws as aws
import pulumi_pulumi as pulumi

primary = pulumi.providers.Aws("primary", region="us-east-2")
secondary = pulumi.providers.Aws("secondary", region="us-west-2")
example = aws.rds.GlobalCluster("example", global_cluster_identifier="example",
opts=ResourceOptions(provider="aws.primary"))
primary_cluster = aws.rds.Cluster("primaryCluster",
    engine_mode="global",
    global_cluster_identifier=example.id,
    opts=ResourceOptions(provider="aws.primary"))
primary_cluster_instance = aws.rds.ClusterInstance("primaryClusterInstance", cluster_identifier=primary_cluster.id,
opts=ResourceOptions(provider="aws.primary"))
secondary_cluster = aws.rds.Cluster("secondaryCluster",
    engine_mode="global",
    global_cluster_identifier=example.id,
    opts=ResourceOptions(provider="aws.secondary",
        depends_on=["aws_rds_cluster_instance.primary"]))
secondary_cluster_instance = aws.rds.ClusterInstance("secondaryClusterInstance", cluster_identifier=secondary_cluster.id,
opts=ResourceOptions(provider="aws.secondary"))

