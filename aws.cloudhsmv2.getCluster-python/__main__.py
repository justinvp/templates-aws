import pulumi
import pulumi_aws as aws

cluster = aws.cloudhsmv2.get_cluster(cluster_id="cluster-testclusterid")

