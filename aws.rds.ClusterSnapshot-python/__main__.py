import pulumi
import pulumi_aws as aws

example = aws.rds.ClusterSnapshot("example",
    db_cluster_identifier=aws_rds_cluster["example"]["id"],
    db_cluster_snapshot_identifier="resourcetestsnapshot1234")

