using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.WafRegional.RegexPatternSet("example", new Aws.WafRegional.RegexPatternSetArgs
        {
            RegexPatternStrings = 
            {
                "one",
                "two",
            },
        });
    }

}

