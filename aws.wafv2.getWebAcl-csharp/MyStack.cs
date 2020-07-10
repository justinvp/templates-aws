using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.WafV2.GetWebAcl.InvokeAsync(new Aws.WafV2.GetWebAclArgs
        {
            Name = "some-web-acl",
            Scope = "REGIONAL",
        }));
    }

}

