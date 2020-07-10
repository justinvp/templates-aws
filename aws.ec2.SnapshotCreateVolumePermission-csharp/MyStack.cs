using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Ebs.Volume("example", new Aws.Ebs.VolumeArgs
        {
            AvailabilityZone = "us-west-2a",
            Size = 40,
        });
        var exampleSnapshot = new Aws.Ebs.Snapshot("exampleSnapshot", new Aws.Ebs.SnapshotArgs
        {
            VolumeId = example.Id,
        });
        var examplePerm = new Aws.Ec2.SnapshotCreateVolumePermission("examplePerm", new Aws.Ec2.SnapshotCreateVolumePermissionArgs
        {
            AccountId = "12345678",
            SnapshotId = exampleSnapshot.Id,
        });
    }

}

