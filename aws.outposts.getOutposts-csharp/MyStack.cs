using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Outposts.GetOutposts.InvokeAsync(new Aws.Outposts.GetOutpostsArgs
        {
            SiteId = data.Aws_outposts_site.Id,
        }));
    }

}

