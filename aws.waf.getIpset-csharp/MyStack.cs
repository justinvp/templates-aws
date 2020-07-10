using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Waf.GetIpset.InvokeAsync(new Aws.Waf.GetIpsetArgs
        {
            Name = "tfWAFIPSet",
        }));
    }

}

