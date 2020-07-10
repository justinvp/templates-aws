using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Outposts.GetOutpost.InvokeAsync(new Aws.Outposts.GetOutpostArgs
        {
            Name = "example",
        }));
    }

}

