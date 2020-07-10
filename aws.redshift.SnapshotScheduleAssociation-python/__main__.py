import pulumi
import pulumi_aws as aws

default_cluster = aws.redshift.Cluster("defaultCluster",
    cluster_identifier="tf-redshift-cluster",
    cluster_type="single-node",
    database_name="mydb",
    master_password="Mustbe8characters",
    master_username="foo",
    node_type="dc1.large")
default_snapshot_schedule = aws.redshift.SnapshotSchedule("defaultSnapshotSchedule",
    definitions=["rate(12 hours)"],
    identifier="tf-redshift-snapshot-schedule")
default_snapshot_schedule_association = aws.redshift.SnapshotScheduleAssociation("defaultSnapshotScheduleAssociation",
    cluster_identifier=default_cluster.id,
    schedule_identifier=default_snapshot_schedule.id)

