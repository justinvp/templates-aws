import pulumi
import pulumi_aws as aws

postgresql = aws.rds.Cluster("postgresql",
    availability_zones=[
        "us-west-2a",
        "us-west-2b",
        "us-west-2c",
    ],
    backup_retention_period=5,
    cluster_identifier="aurora-cluster-demo",
    database_name="mydb",
    engine="aurora-postgresql",
    master_password="bar",
    master_username="foo",
    preferred_backup_window="07:00-09:00")

