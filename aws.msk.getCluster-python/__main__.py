import pulumi
import pulumi_aws as aws

example = aws.msk.get_cluster(cluster_name="example")

