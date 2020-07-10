using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Outposts.GetSite.InvokeAsync(new Aws.Outposts.GetSiteArgs
        {
            Name = "example",
        }));
    }

}

