using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.WafV2.GetIpSet.InvokeAsync(new Aws.WafV2.GetIpSetArgs
        {
            Name = "some-ip-set",
            Scope = "REGIONAL",
        }));
    }

}

