using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.Waf.GetRateBasedRule.InvokeAsync(new Aws.Waf.GetRateBasedRuleArgs
        {
            Name = "tfWAFRateBasedRule",
        }));
    }

}

