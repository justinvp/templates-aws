using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleRegexPatternSet = new Aws.WafRegional.RegexPatternSet("exampleRegexPatternSet", new Aws.WafRegional.RegexPatternSetArgs
        {
            RegexPatternStrings = 
            {
                "one",
                "two",
            },
        });
        var exampleRegexMatchSet = new Aws.WafRegional.RegexMatchSet("exampleRegexMatchSet", new Aws.WafRegional.RegexMatchSetArgs
        {
            RegexMatchTuples = 
            {
                new Aws.WafRegional.Inputs.RegexMatchSetRegexMatchTupleArgs
                {
                    FieldToMatch = new Aws.WafRegional.Inputs.RegexMatchSetRegexMatchTupleFieldToMatchArgs
                    {
                        Data = "User-Agent",
                        Type = "HEADER",
                    },
                    RegexPatternSetId = exampleRegexPatternSet.Id,
                    TextTransformation = "NONE",
                },
            },
        });
    }

}

