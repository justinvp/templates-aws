import pulumi
import pulumi_aws as aws

example = aws.ebs.Volume("example",
    availability_zone="us-west-2a",
    size=40)
example_snapshot = aws.ebs.Snapshot("exampleSnapshot", volume_id=example.id)
example_perm = aws.ec2.SnapshotCreateVolumePermission("examplePerm",
    account_id="12345678",
    snapshot_id=example_snapshot.id)

