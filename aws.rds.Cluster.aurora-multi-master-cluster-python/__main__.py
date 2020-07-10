import pulumi
import pulumi_aws as aws

example = aws.rds.Cluster("example",
    cluster_identifier="example",
    db_subnet_group_name=aws_db_subnet_group["example"]["name"],
    engine_mode="multimaster",
    master_password="barbarbarbar",
    master_username="foo",
    skip_final_snapshot=True)

