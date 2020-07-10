import pulumi
import pulumi_aws as aws

example = aws.ebs.Volume("example",
    availability_zone="us-west-2a",
    size=40,
    tags={
        "Name": "HelloWorld",
    })
example_snapshot = aws.ebs.Snapshot("exampleSnapshot",
    tags={
        "Name": "HelloWorld_snap",
    },
    volume_id=example.id)
example_copy = aws.ebs.SnapshotCopy("exampleCopy",
    source_region="us-west-2",
    source_snapshot_id=example_snapshot.id,
    tags={
        "Name": "HelloWorld_copy_snap",
    })

