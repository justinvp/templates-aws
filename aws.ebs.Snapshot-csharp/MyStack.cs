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
            Tags = 
            {
                { "Name", "HelloWorld" },
            },
        });
        var exampleSnapshot = new Aws.Ebs.Snapshot("exampleSnapshot", new Aws.Ebs.SnapshotArgs
        {
            Tags = 
            {
                { "Name", "HelloWorld_snap" },
            },
            VolumeId = example.Id,
        });
    }

}

