import pulumi
import pulumi_aws as aws

config = pulumi.Config()
cluster_name = config.get("clusterName")
if cluster_name is None:
    cluster_name = "example"
example_cluster = aws.eks.Cluster("exampleCluster", enabled_cluster_log_types=[
    "api",
    "audit",
],
opts=ResourceOptions(depends_on=["aws_cloudwatch_log_group.example"]))
example_log_group = aws.cloudwatch.LogGroup("exampleLogGroup", retention_in_days=7)

