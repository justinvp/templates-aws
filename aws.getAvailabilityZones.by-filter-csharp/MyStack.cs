using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.GetAvailabilityZones.InvokeAsync(new Aws.GetAvailabilityZonesArgs
        {
            AllAvailabilityZones = true,
            Filters = 
            {
                new Aws.Inputs.GetAvailabilityZonesFilterArgs
                {
                    Name = "opt-in-status",
                    Values = 
                    {
                        "not-opted-in",
                        "opted-in",
                    },
                },
            },
        }));
    }

}

