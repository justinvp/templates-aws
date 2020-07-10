import pulumi
import pulumi_aws as aws

replica = aws.elasticache.Cluster("replica", replication_group_id=aws_elasticache_replication_group["example"]["id"])

