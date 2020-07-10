import pulumi
import pulumi_aws as aws

docdb = aws.docdb.Cluster("docdb",
    backup_retention_period=5,
    cluster_identifier="my-docdb-cluster",
    engine="docdb",
    master_password="mustbeeightchars",
    master_username="foo",
    preferred_backup_window="07:00-09:00",
    skip_final_snapshot=True)

