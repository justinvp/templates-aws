using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Pinpoint.App("example", new Aws.Pinpoint.AppArgs
        {
            Limits = new Aws.Pinpoint.Inputs.AppLimitsArgs
            {
                MaximumDuration = 600,
            },
            QuietTime = new Aws.Pinpoint.Inputs.AppQuietTimeArgs
            {
                End = "06:00",
                Start = "00:00",
            },
        });
    }

}

