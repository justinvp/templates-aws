using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.WafV2.RegexPatternSet("example", new Aws.WafV2.RegexPatternSetArgs
        {
            Description = "Example regex pattern set",
            RegularExpressions = 
            {
                new Aws.WafV2.Inputs.RegexPatternSetRegularExpressionArgs
                {
                    RegexString = "one",
                },
                new Aws.WafV2.Inputs.RegexPatternSetRegularExpressionArgs
                {
                    RegexString = "two",
                },
            },
            Scope = "REGIONAL",
            Tags = 
            {
                { "Tag1", "Value1" },
                { "Tag2", "Value2" },
            },
        });
    }

}

