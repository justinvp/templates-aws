using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.WafRegional.GetRateBasedMod.InvokeAsync(new Aws.WafRegional.GetRateBasedModArgs
        {
            Name = "tfWAFRegionalRateBasedRule",
        }));
    }

}

