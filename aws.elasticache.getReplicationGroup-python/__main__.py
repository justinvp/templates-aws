import pulumi
import pulumi_aws as aws

bar = aws.elasticache.get_replication_group(replication_group_id="example")

