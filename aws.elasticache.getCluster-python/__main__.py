import pulumi
import pulumi_aws as aws

my_cluster = aws.elasticache.get_cluster(cluster_id="my-cluster-id")

