using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.WafRegional.GetIpset.InvokeAsync(new Aws.WafRegional.GetIpsetArgs
        {
            Name = "tfWAFRegionalIPSet",
        }));
    }

}

