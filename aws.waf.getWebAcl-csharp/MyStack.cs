using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Waf.GetWebAcl.InvokeAsync(new Aws.Waf.GetWebAclArgs
        {
            Name = "tfWAFWebACL",
        }));
    }

}

