using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = Output.Create(Aws.WafV2.GetRegexPatternSet.InvokeAsync(new Aws.WafV2.GetRegexPatternSetArgs
        {
            Name = "some-regex-pattern-set",
            Scope = "REGIONAL",
        }));
    }

}

