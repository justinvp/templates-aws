using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.WafRegional.GetRule.InvokeAsync(new Aws.WafRegional.GetRuleArgs
        {
            Name = "tfWAFRegionalRule",
        }));
    }

}

