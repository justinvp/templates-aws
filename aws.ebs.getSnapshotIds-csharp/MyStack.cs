using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var ebsVolumes = Output.Create(Aws.Ebs.GetSnapshotIds.InvokeAsync(new Aws.Ebs.GetSnapshotIdsArgs
        {
            Filters = 
            {
                new Aws.Ebs.Inputs.GetSnapshotIdsFilterArgs
                {
                    Name = "volume-size",
                    Values = 
                    {
                        "40",
                    },
                },
                new Aws.Ebs.Inputs.GetSnapshotIdsFilterArgs
                {
                    Name = "tag:Name",
                    Values = 
                    {
                        "Example",
                    },
                },
            },
            Owners = 
            {
                "self",
            },
        }));
    }

}

