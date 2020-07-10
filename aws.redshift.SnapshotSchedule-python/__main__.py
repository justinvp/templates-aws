import pulumi
import pulumi_aws as aws

default = aws.redshift.SnapshotSchedule("default",
    definitions=["rate(12 hours)"],
    identifier="tf-redshift-snapshot-schedule")

