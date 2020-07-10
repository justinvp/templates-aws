using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.WafRegional.GetWebAcl.InvokeAsync(new Aws.WafRegional.GetWebAclArgs
        {
            Name = "tfWAFRegionalWebACL",
        }));
    }

}

