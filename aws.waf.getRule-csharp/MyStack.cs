using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Waf.GetRule.InvokeAsync(new Aws.Waf.GetRuleArgs
        {
            Name = "tfWAFRule",
        }));
    }

}

