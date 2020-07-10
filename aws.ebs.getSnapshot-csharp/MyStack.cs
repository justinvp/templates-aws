using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var ebsVolume = Output.Create(Aws.Ebs.GetSnapshot.InvokeAsync(new Aws.Ebs.GetSnapshotArgs
        {
            Filters = 
            {
                new Aws.Ebs.Inputs.GetSnapshotFilterArgs
                {
                    Name = "volume-size",
                    Values = 
                    {
                        "40",
                    },
                },
                new Aws.Ebs.Inputs.GetSnapshotFilterArgs
                {
                    Name = "tag:Name",
                    Values = 
                    {
                        "Example",
                    },
                },
            },
            MostRecent = true,
            Owners = 
            {
                "self",
            },
        }));
    }

}

