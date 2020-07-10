using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleRegexPatternSet = new Aws.Waf.RegexPatternSet("exampleRegexPatternSet", new Aws.Waf.RegexPatternSetArgs
        {
            RegexPatternStrings = 
            {
                "one",
                "two",
            },
        });
        var exampleRegexMatchSet = new Aws.Waf.RegexMatchSet("exampleRegexMatchSet", new Aws.Waf.RegexMatchSetArgs
        {
            RegexMatchTuples = 
            {
                new Aws.Waf.Inputs.RegexMatchSetRegexMatchTupleArgs
                {
                    FieldToMatch = new Aws.Waf.Inputs.RegexMatchSetRegexMatchTupleFieldToMatchArgs
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

