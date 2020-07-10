import pulumi
import pulumi_aws as aws

test_snapshot_copy_grant = aws.redshift.SnapshotCopyGrant("testSnapshotCopyGrant", snapshot_copy_grant_name="my-grant")
test_cluster = aws.redshift.Cluster("testCluster", snapshot_copy={
    "destinationRegion": "us-east-2",
    "grantName": test_snapshot_copy_grant.snapshot_copy_grant_name,
})

