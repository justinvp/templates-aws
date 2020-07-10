using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var ebsVolume = Output.Create(Aws.Ebs.GetVolume.InvokeAsync(new Aws.Ebs.GetVolumeArgs
        {
            Filters = 
            {
                new Aws.Ebs.Inputs.GetVolumeFilterArgs
                {
                    Name = "volume-type",
                    Values = 
                    {
                        "gp2",
                    },
                },
                new Aws.Ebs.Inputs.GetVolumeFilterArgs
                {
                    Name = "tag:Name",
                    Values = 
                    {
                        "Example",
                    },
                },
            },
            MostRecent = true,
        }));
    }

}

