import pulumi
import pulumi_aws as aws

cluster_name = aws.rds.get_cluster(cluster_identifier="clusterName")

