using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.GameLift.Alias("example", new Aws.GameLift.AliasArgs
        {
            Description = "Example Description",
            RoutingStrategy = new Aws.GameLift.Inputs.AliasRoutingStrategyArgs
            {
                Message = "Example Message",
                Type = "TERMINAL",
            },
        });
    }

}

